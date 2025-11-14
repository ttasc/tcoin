## Vai trò các nút trong bitcoin

1. **Miner**

    Các nút như vậy được vận hành trên phần cứng mạnh mẽ hoặc chuyên dụng (như ASIC), và mục tiêu duy nhất của chúng là đào các block mới càng nhanh càng tốt. Thợ đào chỉ có thể làm việc trong các blockchain sử dụng Proof-of-Work (Bằng chứng Công việc), vì đào thực chất là giải các câu đố PoW. Ví dụ, trong các blockchain Proof-of-Stake (Bằng chứng Cổ phần), không có hoạt động đào.

2. **Fullnode**

    Các nút này xác thực các block được khai thác bởi thợ đào và xác minh giao dịch. Để làm được điều này, chúng phải có toàn bộ bản sao của blockchain. Ngoài ra, các nút này còn thực hiện các hoạt động định tuyến, chẳng hạn như giúp các nút khác tìm thấy nhau.

    Mạng lưới cần có nhiều nút đầy đủ vì chính những nút này sẽ đưa ra quyết định: chúng quyết định xem một block hoặc giao dịch có hợp lệ hay không.

3. **SPV (Simplified Payment Verification)**

    SPV là viết tắt của Simplified Payment Verification (Xác minh Thanh toán Đơn giản). Các nút này không lưu trữ toàn bộ bản sao blockchain, nhưng vẫn có thể xác minh các giao dịch (không phải tất cả, mà chỉ một tập hợp con, ví dụ, các giao dịch được gửi đến một địa chỉ cụ thể). Một nút SPV phụ thuộc vào một nút đầy đủ để lấy dữ liệu, và có thể có nhiều nút SPV được kết nối với một nút đầy đủ. SPV giúp các ứng dụng ví trở nên khả thi: người dùng không cần tải xuống toàn bộ blockchain, nhưng vẫn có thể xác minh các giao dịch của mình.

## Đơn giản hóa mạng

Để triển khai mạng lưới trong blockchain, ta cần đơn giản hóa một số thứ. Vấn đề là ta không có nhiều máy tính để mô phỏng một mạng lưới với nhiều nút. Ta sẽ sử dụng Docker để giải quyết vấn đề này, các máy tính (ở đây là các container) sẽ giao tiếp ở cổng 3000.

## Triển khai thực tế

Điều gì sẽ xảy ra khi bạn tải xuống, chẳng hạn như Bitcoin Core và chạy nó lần đầu tiên? Nó phải kết nối với một số nút để tải xuống trạng thái mới nhất của blockchain. Thử nghĩ đến việc máy tính của bạn làm sao biết được một máy tính nào đó khác có lưu trạng thái blockchain để yêu cầu tải xuống? Đó là nút nào?

Việc mã hóa cứng địa chỉ nút trong Bitcoin Core sẽ là một sai lầm: nút có thể bị tấn công hoặc bị tắt, dẫn đến việc các nút mới không thể tham gia mạng. Thay vào đó, trong Bitcoin Core, có các hạt giống DNS được hardcode. Chúng không phải là nút, mà là máy chủ DNS biết địa chỉ của một số nút. Khi bạn khởi động một Bitcoin Core sạch, nó sẽ kết nối với một trong các hạt giống và nhận được danh sách các nút đầy đủ, sau đó nó sẽ tải xuống blockchain từ đó.

Tuy nhiên, trong quá trình triển khai, ta sẽ tập trung hóa. Ta sẽ có ba nút:

1. Nút trung tâm. Đây là nút mà tất cả các nút khác sẽ kết nối tới và là nút sẽ gửi dữ liệu giữa các nút khác. Có thể xem như nút này có vai trò là **Fullnode**
2. **Miner node**. Nút này sẽ lưu trữ các giao dịch mới trong mempool và khi có đủ giao dịch, nó sẽ khai thác một block mới.
3. Một nút ví. Nút này sẽ được sử dụng để gửi coin giữa các ví. Tuy nhiên, không giống như các nút **SPV**, nó sẽ lưu trữ một bản sao đầy đủ của blockchain.

> Ta sẽ hardcode địa chỉ của nút trung tâm để tất cả các nút đều có thể liên lạc với nút trung tâm.

## Demo
### Quy ước cách gọi
Trước tiên ta cần quy ước cách gọi để dễ dàng quan sát và hiểu vai trò của từng nút trong demo so với bitcoin:
- **Host**: Là máy tính vật lý có vai trò trung gian để phân phát tài nguyên động thay vì hardcode như trong bitcoin
- **Fullnode**: có vai trò như một nút đầy đủ, và ở trong demo này chúng còn có vai trò là nút trung tâm để tiếp nhận và phân phát dữ liệu
- **Miner**: có vai trò là nút khai thác
- **Wallet**: là các nút ví có vai trò thực hiện giao dịch, nếu bạn dùng tcoin để giao dịch, thiết bị của bạn là một nút Wallet

### Kịch bản
Ta sẽ triển khai tình huống sau:
1. **Fullnode** tạo ra một blockchain.
2. **Wallets** kết nối với nó và tải xuống blockchain.
3. **Miners** kết nối với nút trung tâm và tải xuống blockchain.
4. **Wallets** tạo ra giao dịch.
5. **Miners** nhận giao dịch và lưu giữ trong memory pool.
6. Khi có đủ giao dịch trong memory pool, **Miners** sẽ bắt đầu khai thác một block mới.
7. Khi một block mới được khai thác, nó sẽ được gửi đến **Fullnode**.
8. **Wallets** đồng bộ hóa với nút trung tâm.
9. Người dùng nút ví (**Wallets**) kiểm tra xem thanh toán của họ đã thành công chưa.

### Triển khai

#### Yêu cầu
- Go & source code (Nếu dựng file nhị phân từ mã nguồn)
- Docker

#### Chuẩn bị
Trước tiên ta sẽ dùng docker để tạo ra 3 container dựa trên ubuntu image và đặt hostname cho chúng lần lượt là: `fullnode`, `miner` và `wallet`.

Ta sẽ dùng lớp mạng mặc định mà docker gán cho các container (bridge driver). Lần lượt log-on vào các container theo thứ tự sẽ được các địa chỉ như sau:

- fullnode: `172.17.0.2`
- miner: `172.17.0.3`
- wallet: `172.17.0.4`


#### Thực hiện

##### Fullnode

- Ở đây, thiết bị Fullnode sẽ là nút tạo ra blockchain và tạo ra block đầu tiên trong tcoin, ta cần một địa chỉ để nút này thu về phần thưởng khi tạo ra block đầu tiên (trong bitcoin, địa chỉ nhận phần thưởng của block đầu tiên thuộc về Satoshi):
    ```bash
    [fullnode]$ tcoin createwallet
    Your new address: 13QLoHmb1QrUeZK4DNDv1DP337rojpCW9t
    ```
- Dùng địa chỉ vừa tạo để gắn vào lệnh tạo blockchain sau:
    ```bash
    [fullnode]$ tcoin createblockchain -address 13QLoHmb1QrUeZK4DNDv1DP337rojpCW9t
    0000d20e55a7fa8a601795c7b153b1dda1f8c8591d60fee5219153caacdc69c8

    Done!
    ```
    Sau bước này, blockchain sẽ chứa một genesis block duy nhất. Ta cần lưu block này và sử dụng nó trong các nút khác. Genesis block đóng vai trò là định danh của blockchain (trong Bitcoin Core, khối genesis được *hardcode*).
- Dùng lệnh sau để in ra blockchain trong cửa sổ terminal:
    ```bash
    tcoin printchain
    ```

##### Host
Do ta không hardcode genesis block nên sẽ dùng máy vật lý để phân phát genesis block cho các nút. Trước tiên ta sẽ copy genesis block từ container ra máy vật lý:
```bash
docker cp fullnode:/root/blockchain_172.17.0.2.db blockchain_genesis.db
```

##### Wallet
Chuyển sang các nút Wallet và tạo ra một số ví để thực hiện giao dịch.
- Tạo 4 địa chỉ ví:
    ```bash
    tcoin createwallet # x4
    ```
- Liệt kê các địa chỉ ví vừa tạo:
    ```bash
    [wallet]$ tcoin listaddresses
    1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8
    1BmsTS98y8VgpCmVbncW4WL2u7fy3W4ZTv
    1LQSLPHThtptRGysbbuAv4M7we7RGV7x4q
    1LcyFhFhR8Exc3ANJtG2QjuUhNSuPEWPb
    ```
> Theo thứ tự được liệt kê bởi lệnh trên, ta sẽ gọi chúng là các ví 1, ví 2, ví 3, ví 4.

##### Fullnode
Chuyển sang Fullnode và thực hiện gửi coin cho ví 1 và ví 2.
- Gửi tiền và đào các block:
    ```bash
    tcoin send -from <address> -to <address> -amount 10 -mine
    ```
    > Ta dùng flag -mine để đào các khối ngay khi tạo ra các giao dịch vì ban đầu không hề có nút khai thác nào trong mạng.

- Gửi cho ví 1 và ví 2 mỗi ví 10 coin:
    ```bash
    [fullnode]$ tcoin send -from 13QLoHmb1QrUeZK4DNDv1DP337rojpCW9t -to 1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8 -amount 10 -mine
    000063505c720728db4afc2c264a868f09d644826e5fd3b8368c14e401614fd6

    Success!


    [fullnode]$ tcoin send -from 13QLoHmb1QrUeZK4DNDv1DP337rojpCW9t -to 1BmsTS98y8VgpCmVbncW4WL2u7fy3W4ZTv -amount 10 -mine
    0000f54936e50d84af9a41eee82b265e2f64e36a014add92d75b466fc56cbfd5

    Success!
    ```

- Khởi chạy nút, lúc này Fullnode đã thực sự là một thành phần của mạng blockchain (nút này phải chạy cho đến hết kịch bản):
    ```bash
    tcoin startnode
    ```

##### Host
Trước khi chuyển sang nút Wallet để gia nhập mạng blockchain, nút Wallet cần biết đâu là blockchain mà nó sẽ gia nhập thông qua genesis block (một lần nữa, genesis block được hardcode!), do đó ta cần copy genesis block từ máy vật lý vào `wallet` container:
```bash
docker cp blockchain_genesis.db wallet:/root/blockchain_172.17.0.4.db
```

##### Wallet
- Tại nút Wallet, khởi chạy nút sẽ bắt đầu tải xuống tất cả các block từ nút trung tâm Fullnode:
    ```bash
    tcoin startnode
    ```

- Kiểm tra số dư của ví 1 và ví 2:
    ```bash
    [wallet]$ tcoin listaddresses
    1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8
    1BmsTS98y8VgpCmVbncW4WL2u7fy3W4ZTv
    1LQSLPHThtptRGysbbuAv4M7we7RGV7x4q
    1LcyFhFhR8Exc3ANJtG2QjuUhNSuPEWPb

    [wallet]$ tcoin getbalance -address 1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8
    Balance of '1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8': 10

    [wallet]$ tcoin getbalance -address 1BmsTS98y8VgpCmVbncW4WL2u7fy3W4ZTv
    Balance of '1BmsTS98y8VgpCmVbncW4WL2u7fy3W4ZTv': 10
    ```
- Chúng ta dĩ nhiên cũng có thể kiểm tra số dư của nút Fullnode bởi vì nút Wallet hiện tại cũng đã chứa chuỗi khối của nó!
    ```bash
    [wallet]$ tcoin getbalance -address 13QLoHmb1QrUeZK4DNDv1DP337rojpCW9t
    Balance of '13QLoHmb1QrUeZK4DNDv1DP337rojpCW9t': 10
    ```

##### Host
Được rồi, đã đến lúc các thợ đào gia nhập đội hình! Tương tự như các nút khác, các thợ đào cũng cần phải biết đâu là blockchain chính thống, chúng ta vẫn phải copy genesis block cho các thợ đào:
```bash
docker cp blockchain_genesis.db miner:/root/blockchain_172.17.0.3.db
```

##### Miner
- Trước tiên các thợ đào cần có một địa chỉ ví để nhận phần thưởng khi đào được các block:
    ```bash
    [miner]$ tcoin createwallet
    Your new address: 1BJiAuYqcChkQCMH8LHWWZT177J4PjuHDw
    ```
- **Bắt đầu đào thôi!** dùng flag `-miner` để chỉ định đây là một nút thợ đào trong mạng, các thợ đào sẽ bắt đầu đào một khối khi có từ 2 giao dịch trở lên, hãy khởi động và đợi các nút Wallet giao dịch:
    ```bash
    tcoin startnode -miner 1BJiAuYqcChkQCMH8LHWWZT177J4PjuHDw
    ```

##### Wallet
Chuyển sang nút Wallet và bắt đầu chuyển tiền nào!
- Gửi `1 coin từ ví 1 vào ví 3` và `1 coin từ ví 2 vào ví 4`:
    ```bash
    [wallet]$ tcoin send -from 1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8 -to 1LQSLPHThtptRGysbbuAv4M7we7RGV7x4q -amount 1
    Success!

    [wallet]$ tcoin send -from 1BmsTS98y8VgpCmVbncW4WL2u7fy3W4ZTv -to 1LcyFhFhR8Exc3ANJtG2QjuUhNSuPEWPb -amount 1
    Success!
    ```

Mỗi giao dịch được thực hiện sẽ gửi về cho nút trung tâm Fullnode. Fullnode tiếp nhận giao dịch và broadcast ID của giao dịch đó trên mạng để thông báo có một giao dịch mới được tạo.

Các thợ đào nhận ID của giao dịch và kiểm tra trong mempool đã có giao dịch này chưa, nếu chưa thì các thợ đào sẽ gửi yêu cầu tải xuống dữ liệu của giao dịch đó và lưu vào mempool.

Khi có đủ từ 2 giao dịch trở lên, thợ đào sẽ gắn các giao dịch vào một block và đào block đó. Sau khi đào được một block, thợ đào sẽ broadcast block này trên mạng để các nút khác cập nhật lại blockchain.

- Ngay sau khi thực hiện giao dịch, các thợ đào đã cật lực đào ra block mới và cập nhật lại blockchain. Nhưng nút Wallet của chúng ta hiện vẫn chưa gia nhập mạng blockchain để cập nhật block mới, hãy thử kiểm tra trước khi gia nhập mạng!
    ```bash
    [wallet]$ tcoin getbalance -address 1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8
    Balance of '1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8': 10
    # Số dư của ví 1 vẫn là 10 coin trong khi đã gửi 1 coin cho ví 3, phải là 9 coin mới đúng!
    ```

- Được rồi, đến lúc gia nhập mạng để biết blockchain đã thay đổi như thế nào:
    ```bash
    tcoin startnode
    ```

- Kiểm tra lại số dư của các ví xem đã hợp lý chưa!
    ```bash
    # Nút Wallet (ví 1)
    [wallet]$ tcoin getbalance -address 1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8
    Balance of '1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8': 9

    # Nút Wallet (ví 2)
    [wallet]$ tcoin getbalance -address 1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8
    Balance of '1BmsTS98y8VgpCmVbncW4WL2u7fy3W4ZTv': 9

    # Nút Wallet (ví 3)
    [wallet]$ tcoin getbalance -address 1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8
    Balance of '1LQSLPHThtptRGysbbuAv4M7we7RGV7x4q': 1

    # Nút Wallet (ví 4)
    [wallet]$ tcoin getbalance -address 1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8
    Balance of '1LcyFhFhR8Exc3ANJtG2QjuUhNSuPEWPb': 1

    # Nút trung tâm Fullnode
    [wallet]$ tcoin getbalance -address 1BYYGAJmGiH8XZs2hHcCkd6ix3yaKPLFN8
    Balance of '13QLoHmb1QrUeZK4DNDv1DP337rojpCW9t': 10
    ```
