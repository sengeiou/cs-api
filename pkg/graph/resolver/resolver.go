package resolver

import (
	"cs-api/config"
	generated "cs-api/dist/graph"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AndySu1021/go-util/exporter"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	zlog "github.com/AndySu1021/go-util/zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
)

type Resolver struct {
	authSvc     iface.IAuthService
	staffSvc    iface.IStaffService
	roomSvc     iface.IRoomService
	messageSvc  iface.IMessageService
	tagSvc      iface.ITagService
	reportSvc   iface.IReportService
	csConfigSvc iface.ICsConfigService
	noticeSvc   iface.INoticeService
	remindSvc   iface.IRemindService
	redis       ifaceTool.IRedis
	storage     iface.IStorage
	config      *config.Config
}

var Module = fx.Options(
	fx.Provide(
		createConfig,
		NewResolver,
	),
	fx.Invoke(
		InitResolver,
	),
)

type Params struct {
	fx.In

	AuthSvc     iface.IAuthService
	StaffSvc    iface.IStaffService
	RoomSvc     iface.IRoomService
	MessageSvc  iface.IMessageService
	TagSvc      iface.ITagService
	ReportSvc   iface.IReportService
	CsConfigSvc iface.ICsConfigService
	NoticeSvc   iface.INoticeService
	RemindSvc   iface.IRemindService
	Redis       ifaceTool.IRedis
	Storage     iface.IStorage
	Config      *config.Config
}

func NewResolver(p Params) *Resolver {
	return &Resolver{
		authSvc:     p.AuthSvc,
		staffSvc:    p.StaffSvc,
		roomSvc:     p.RoomSvc,
		messageSvc:  p.MessageSvc,
		tagSvc:      p.TagSvc,
		reportSvc:   p.ReportSvc,
		csConfigSvc: p.CsConfigSvc,
		noticeSvc:   p.NoticeSvc,
		remindSvc:   p.RemindSvc,
		redis:       p.Redis,
		storage:     p.Storage,
		config:      p.Config,
	}
}

func createConfig(r *Resolver) generated.Config {
	c := generated.Config{
		Resolvers:  r,
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}

	return c
}

func InitResolver(logCfg *zlog.Config, engine *gin.Engine, cfg generated.Config, authSvc iface.IAuthService) error {
	gqlSvc := handler.New(generated.NewExecutableSchema(cfg))

	// Set transport policy
	gqlSvc.AddTransport(transport.POST{})
	gqlSvc.AddTransport(transport.MultipartForm{})
	gqlSvc.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			HandshakeTimeout: 15 * time.Second,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
				log.Error().Msgf("ws error: %s", reason)
			},
		},
	})

	// Set cache
	gqlSvc.SetQueryCache(lru.New(1000))

	// Set middleware
	gqlSvc.Use(extension.FixedComplexityLimit(500))
	gqlSvc.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	//gqlSvc.AroundResponses(graph.GQLResponseLog(&graph.Config{}))
	//gqlSvc.AroundResponses(auditLogSvc.RecordAuditLogForGraphql)
	//gqlSvc.SetErrorPresenter(errors.GQLErrorPresenter)
	//gqlSvc.SetRecoverFunc(graph.GQLRecoverFunc)

	exporter.GraphQLRegister()
	tracer := exporter.GraphQLTracer{}
	gqlSvc.Use(tracer)

	engine.Any("/graph/query/staff", authSvc.SetClientInfo(pkg.ClientTypeStaff), gin.WrapH(gqlSvc))
	engine.Any("/graph/query/member", authSvc.SetClientInfo(pkg.ClientTypeMember), gin.WrapH(gqlSvc))

	if logCfg.Environment != "prod" {
		gqlSvc.Use(extension.Introspection{})
		playGround := playground.Handler("GraphQL Playground", "/graph/query")
		engine.Any("/graph/playground", gin.WrapH(playGround))
	}

	return nil
}
