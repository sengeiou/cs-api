# 客服系統

## websocket 壓測結果

## 系統問題
1. 由於使用 map 做緩存，紀錄房間號碼與其對應的 connection，高併發的情況下會造成 concurrent map read and map write 的文題
   解法：透過 sync.Map 解決
2. 