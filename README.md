# Reference
1. https://learnblockchain.cn/article/5782
2. https://docs.unisat.io/dev/unisat-developer-service

# Steps
铭文交易的构建过程包含以下步骤：
    准备铸造的内容和可用的 UTXO。
    构建铭文交易，包括 commit_tx 和 reveal_tx。
    创建一个 SatPoint，也就是铭文所在的 OutPoint。
    检查这个 OutPoint 是否已经铭刻过。
    获取一个 secp256k1 的密钥对，用于构建 reveal_tx。
    构建 reveal 脚本，将铭文内容写入其中。
    构建 Taproot 花费脚本（taproot spend script），将 reveal 脚本添加到脚本的 leaf 上。
    从花费脚本中获取 tweaked key。
    计算 reveal 交易的费用。
    构建 commit 交易。
    用 commit 交易构建 reveal 交易。
    对 reveal 交易进行数量检查。
    计算签名哈希。
    将签名哈希推入 witness 数据。
    用钱包对 commit_tx 进行签名，发送 commit_tx 和 reveal_tx 到比特币网络。
    铭文交易的构建过程非常复杂，需要涉及到多个参数和计算，但是这个过程也非常精妙和有趣。