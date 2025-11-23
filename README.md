## Giới thiệu

Blockchain là một trong những công nghệ mang tính cách mạng nhất của thế kỷ 21, vẫn đang trong quá trình hoàn thiện và tiềm năng của nó vẫn chưa được khai thác hết. Về bản chất, blockchain chỉ là một cơ sở dữ liệu phân tán chứa các bản ghi. Nhưng điều làm nên sự độc đáo của nó là nó không phải là một cơ sở dữ liệu riêng tư mà là một cơ sở dữ liệu công khai, tức là mọi người sử dụng nó đều có bản sao đầy đủ hoặc một phần của nó. Và một bản ghi mới chỉ có thể được thêm vào khi có sự đồng ý của những người quản lý cơ sở dữ liệu khác. Hơn nữa, chính blockchain đã tạo ra tiền điện tử và hợp đồng thông minh.

Chúng ta sẽ xây dựng một loại tiền điện tử đơn giản dựa trên nền tảng blockchain.

## Block

Bắt đầu với phần "block" của "blockchain". Trong blockchain, các block lưu trữ thông tin có giá trị. Ví dụ, khối bitcoin lưu trữ các giao dịch, bản chất của bất kỳ loại tiền điện tử nào. Bên cạnh đó, một khối còn chứa một số thông tin kỹ thuật, chẳng hạn như phiên bản, dấu thời gian hiện tại và mã băm của khối trước đó.

Đối với dự án này, chúng ta sẽ không triển khai block theo đúng mô tả kỹ thuật của blockchain hay bitcoin mà chỉ chứa các thông tin quan trọng.

Ban đầu, một cách sơ khai nhất - nó sẽ trông như thế này:

```go
# Go programing language
type Block struct {
    Timestamp       int64   # Mốc thời gian khi block được tạo
    Data            []byte  # Thông tin có giá trị trong block
    PrevBlockHash   []byte  # Mã băm của block trước đó
    Hash            []byte  # Mã băm của block
}
```
> Trong đặc tả bitcoin `Timestamp`, `PrevBlockHash` và `Hash` là các **Block Headers**, tách biệt với **Transactions** ở đây chính là `Data`.

## Blockchain

Chuỗi khối - blockchain về bản chất chỉ là một cơ sở dữ liệu với cấu trúc nhất định: nó là một danh sách được sắp xếp theo thứ tự và liên kết ngược, nghĩa là các khối được liên kết với khối trước đó. Cấu trúc này cho phép nhanh chóng lấy khối mới nhất trong chuỗi và (một cách hiệu quả) lấy khối theo hàm băm của nó.

#### Genesis block

Để thêm một khối mới, chúng ta cần một khối trước đó, nhưng blockchain của chúng ta ban đầu thì không có khối nào cả! Vậy nên, trong bất kỳ blockchain nào, phải có ít nhất một khối, và khối đó - khối đầu tiên trong chuỗi, được gọi là khối khởi tạo (genesis block). Khối genesis còn được coi như khối định danh cho cả chuỗi khối.

> ##### Facts:
> Khối genesis của bitcoin luôn được **mã hóa cứng** (hardcode) vào các phần mềm sử dụng chuỗi khối của bitcoin.

## Proof-of-Work
Ý tưởng cốt lõi của blockchain là người ta phải bỏ công sức để đưa dữ liệu vào đó. Chính công sức này làm cho blockchain an toàn và nhất quán. Hơn nữa, có phần thưởng được trả cho công sức này (đây là cách mọi người nhận được tiền khi khai thác).

Cơ chế này rất giống với cơ chế ngoài đời thực: người ta phải làm việc chăm chỉ để nhận được phần thưởng và duy trì cuộc sống. Trong blockchain, một số người tham gia (thợ đào) của mạng lưới làm việc để duy trì mạng lưới, thêm các khối mới vào đó và nhận phần thưởng cho công việc của họ. Nhờ công sức của họ, một khối được tích hợp vào blockchain một cách an toàn, giúp duy trì tính ổn định của toàn bộ cơ sở dữ liệu blockchain. Cần lưu ý rằng, người hoàn thành công việc phải chứng minh điều này.

Toàn bộ cơ chế "do hard work and prove" này được gọi là *bằng chứng công việc* (**proof-of-work**). Nó khó khăn vì đòi hỏi rất nhiều sức mạnh tính toán: ngay cả máy tính hiệu năng cao cũng không thể thực hiện nhanh chóng. Hơn nữa, độ khó của công việc này tăng dần theo thời gian để duy trì tốc độ khối mới ở mức khoảng 6 khối mỗi giờ. Trong Bitcoin, mục tiêu của công việc này là tìm ra một hàm băm cho một khối, đáp ứng một số yêu cầu. Và chính hàm băm này đóng vai trò là bằng chứng. Do đó, việc tìm ra bằng chứng thực sự là công việc.

Một điều cuối cùng cần lưu ý. Các thuật toán Proof-of-Work phải đáp ứng một yêu cầu: thực hiện công việc thì khó, nhưng xác minh bằng chứng thì dễ.

## Hashing

Tiếp theo ta sẽ thảo luận về hashing - băm.

Băm là một quá trình thu thập giá trị băm cho dữ liệu được chỉ định. Giá trị băm là một biểu diễn duy nhất của dữ liệu mà nó được tính toán. Hàm băm là một hàm lấy dữ liệu có kích thước tùy ý và tạo ra giá trị băm có kích thước cố định. Dưới đây là một số tính năng chính của băm:

1. Dữ liệu gốc không thể được khôi phục từ hàm băm. Do đó, băm không phải là mã hóa.
2. Một dữ liệu nhất định chỉ có thể tạo ra một giá trị băm và giá trị băm đó là duy nhất.
3. Chỉ cần thay đổi một byte trong dữ liệu đầu vào cũng sẽ tạo ra một giá trị băm hoàn toàn khác.

![](https://jeiwan.net/images/hashing-example.png)

Trong blockchain, hàm băm được sử dụng để đảm bảo tính nhất quán của một khối. Dữ liệu đầu vào cho thuật toán băm chứa hàm băm của khối trước đó, do đó việc sửa đổi một khối trong chuỗi là bất khả thi (hoặc ít nhất là khá khó khăn): người ta phải tính toán lại hàm băm của khối đó và hàm băm của tất cả các khối sau nó.

## Hashcash

> ##### Hashcash - Wikipedia:
> *Hashcash là một hệ thống bằng chứng công việc (proof-of-work) được sử dụng để hạn chế email rác và các cuộc tấn công từ chối dịch vụ. Hashcash được Adam Back đề xuất vào năm 1997. Và được mô tả chính thức hơn trong bài báo năm 2002 của Back "Hashcash – A Denial of Service Counter-Measure". Trong Hashcash, máy khách phải nối một số ngẫu nhiên với một chuỗi nhiều lần và băm chuỗi mới này. Sau đó, máy khách phải làm như vậy nhiều lần cho đến khi tìm được một chuỗi băm bắt đầu bằng một số lượng số 0 nhất định.*

Qua đoạn thông tin từ wiki trên, ngay lập tức ta có thể kết luận đây là một thuật toán dạng **brute force**: bạn thay đổi bộ đếm, tính toán mã băm mới, kiểm tra nó, tiếp tục tăng bộ đếm nếu sai, tính toán băm, ...v.v. Đó là lý do tại sao nó tốn kém về mặt tính toán.

Cùng xem xét kỹ hơn các yêu cầu mà một hàm băm phải đáp ứng. Trong triển khai gốc Hashcash, yêu cầu này có thể hiểu đơn giản là "20 bit đầu tiên của hàm băm phải là số 0". Trong Bitcoin, yêu cầu này được điều chỉnh theo thời gian, bởi vì theo thiết kế, một khối phải được tạo ra sau mỗi 10 phút, nhưng sức mạnh tính toán sẽ tăng theo thời gian và ngày càng có nhiều thợ đào tham gia mạng lưới.

Để dễ hình dung, hãy dùng dữ liệu từ ví dụ trước ("I like donuts") và gắn thêm vào đó một giá trị để mã băm của nó sẽ bắt đầu bằng 3 số 0:

![](https://jeiwan.net/images/hashcash-example.png)

`ca07ca` là giá trị thập lục phân của bộ đếm, tức là 13240266 trong hệ thập phân. Và con số ngẫu nhiên mà ta đã tìm được này gọi là `nonce`.

```go
type Block struct {
    Timestamp     int64
    Data          []byte
    PrevBlockHash []byte
    Hash          []byte
    Nonce         int
}
```
Giờ đây `nonce` phải được lưu thành một thuộc tính của Block vì nó cần thiết để xác minh một "bằng chứng".

## Giao dịch - Transactions

Giao dịch là cốt lõi của Bitcoin, và mục đích duy nhất của blockchain là lưu trữ giao dịch một cách an toàn và đáng tin cậy, để không ai có thể sửa đổi chúng sau khi chúng được tạo ra.

### Giao dịch trong Bitcoin

Giao dịch là sự kết hợp giữa **inputs** và **outputs**.

```go
type Transaction struct {
    ID   []byte
    Vin  []TXInput
    Vout []TXOutput
}
```
**inputs** của một giao dịch mới tham chiếu đến **outputs** của một giao dịch trước đó. **outputs** là nơi lưu trữ coins.

![](https://jeiwan.net/images/transactions-diagram.png)

Lưu ý:

1. Có những **outputs** không liên kết đến bất kỳ **inputs** nào.
2. Trong một giao dịch, **inputs** có thể tham chiếu đến **outputs** của nhiều giao dịch.
3. **inputs** bắt buộc phải tham chiếu đến ít nhất một **outputs**.

### Outputs

```go
type TXOutput struct {
    Value        int
    ScriptPubKey string
}
```
**outputs** chính là nơi lưu trữ "coins" (trường `value`). Và "lưu trữ" ở đây có nghĩa là khóa chúng bằng một **câu đố**, được lưu trữ trong `ScriptPubKey`. Trong thực tế, Bitcoin sử dụng một ngôn ngữ kịch bản gọi là *Script* , được dùng để xác định logic khóa và mở khóa **outputs**. Ngôn ngữ này khá thô sơ (điều này được tạo ra một cách có chủ đích để tránh các vụ hack và lạm dụng tiềm ẩn), nhưng chúng ta sẽ không thảo luận chi tiết và cũng không triển khai nó trong dự án.

> Trong Bitcoin, trường giá trị lưu trữ số satoshi , chứ không phải số BTC. Một satoshi bằng một phần trăm triệu của một bitcoin (0,00000001 BTC), do đó đây là đơn vị tiền tệ nhỏ nhất trong Bitcoin.

Một điều quan trọng về **outputs** là chúng **không thể chia nhỏ**, nghĩa là ta không thể tham chiếu một phần giá trị của nó. Khi một đầu ra được tham chiếu trong một giao dịch mới, nó sẽ được chi tiêu như một tổng thể. Và nếu giá trị của nó lớn hơn giá trị yêu cầu, một khoản tiền thừa sẽ được tạo ra và gửi lại cho người gửi. Điều này tương tự như tình huống thực tế khi bạn trả tiền, chẳng hạn, một tờ tiền 5k cho một thứ có giá 1k và nhận lại tiền thừa là 4k.

### Inputs

```go
type TXInput struct {
    Txid      []byte
    Vout      int
    ScriptSig string
}
```

Như đã nói trước đó, **inputs** tham chiếu đến **outputs** trước đó: `Txid` lưu trữ ID của giao dịch hiện tại và `Vout` lưu trữ index của **outputs** trong giao dịch. `ScriptSig` là một tập lệnh cung cấp dữ liệu để sử dụng trong **outputs** `ScriptPubKey`. Nếu dữ liệu chính xác, **outputs** có thể được mở khóa và giá trị của nó có thể được sử dụng để tạo ra các **outputs** mới; nếu không chính xác, **outputs** không thể được tham chiếu trong các **inputs** từ các giao dịch mới. Đây là cơ chế đảm bảo người dùng không thể chi tiêu coins của người khác.

## Quả trứng có trước con gà

Trong Bitcoin, quả trứng có trước con gà. Logic **inputs** tham chiếu **outputs** chính là tình huống kinh điển "con gà có trước hay quả trứng có trước". Và trong Bitcoin, đầu ra có trước đầu vào.

Khi một thợ đào bắt đầu khai thác một khối, họ sẽ thêm một giao dịch **Coinbase** vào khối đó. Giao dịch Coinbase là một loại giao dịch đặc biệt, không yêu cầu **outputs** đã tồn tại trước đó. Nó tạo ra **outputs** (đồng nghĩa với việc tạo ra coins) từ hư không. Đây là phần thưởng mà thợ đào nhận được khi khai thác các khối mới.

Như bạn đã biết, có khối genesis ở đầu blockchain. Chính khối này tạo ra đầu ra đầu tiên trong blockchain. Và không cần đầu ra trước đó vì không có giao dịch nào trước đó và cũng không có đầu ra nào như vậy.

> ##### Facts:
> Trong Bitcoin, giao dịch coinbase đầu tiên chứa thông điệp sau: "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks".

Trong triển khai của chúng ta, `subsidy` là số tiền thưởng. Trong Bitcoin, con số này không được lưu trữ ở bất kỳ đâu và chỉ được tính toán dựa trên tổng số khối: số khối được chia cho 210000. Khai thác khối genesis đã tạo ra 50 BTC và mỗi 210000 khối, phần thưởng sẽ giảm một nửa. Trong quá trình triển khai, chúng ta sẽ lưu trữ phần thưởng dưới dạng *hằng số*.

## Đầu ra của giao dịch nhưng chưa sử dụng - Unspent Transaction Outputs

Chúng ta cần tìm tất cả các **outputs** giao dịch chưa sử dụng (Unspent Transaction Outputs - UTXO). "Chưa sử dụng" nghĩa là các **outputs** này chưa được tham chiếu trong bất kỳ **inputs** nào.

Và tất nhiên khi ta kiểm tra số dư, ta không cần tất cả các **outputs**, mà chỉ cần các **outputs** có thể mở khóa bằng khóa mà chúng ta giữ.

Tóm lại, **outputs** chưa sử dụng là **outputs** chưa được tham chiếu trong bất kỳ **inputs** nào trong blockchain và số dư của một ví là tổng giá trị từ các **outputs** chưa sử dụng được khóa bằng khóa của ví đó.

## Gửi coins như thế nào?

Để gửi một số coins cho ai đó chúng ta cần tạo một giao dịch mới, đặt nó vào một khối và đào khối đó.

Giả sử ví A muốn gửi `n` coins cho ví B:
1. Đầu tiên, phải tìm trong blockchain một lượng coin lớn hơn hoặc bằng `n` chưa chi tiêu thuộc về ví A và danh sách các index của các **outputs** này.
2. Từ các **outputs** chưa chi vừa tìm, ta tạo một danh sách các **inputs** tương ứng tham chiếu đến các **outputs** này.
3. Tiếp theo, tạo một **outputs** và khóa nó với địa chỉ của người nhận. Nếu số coins tổng thể mà ta tìm được lớn hơn số tiền cần gửi, khi này phải tạo thêm một **outputs** nữa có giá trị bằng phần tiền thừa và khóa nó với địa chỉ của người gửi.
4. Khi đã có đủ **inputs** và **outputs** hợp lệ, một **Transaction** sẽ được tạo ra.
5. Cuối cùng, sau khi đã có đủ dữ liệu cần thiết, việc "đào" sẽ được tiến hành sau đó.

Gửi coin nghĩa là tạo một giao dịch và thêm nó vào blockchain thông qua việc đào một khối. Nhưng Bitcoin không thực hiện việc này ngay lập tức. Thay vào đó, nó đưa tất cả các giao dịch mới vào một nhóm bộ nhớ (mempool), và khi một thợ đào sẵn sàng đào một khối, nó sẽ lấy tất cả các giao dịch từ nhóm bộ nhớ và tạo ra một khối ứng viên. Các giao dịch chỉ được xác nhận khi một khối chứa chúng được đào và thêm vào blockchain.

## Phần tiếp theo đang cập nhật...


















---
## Demo

### 1. Vai trò các nút trong bitcoin

1. **Miner**

    Các nút như vậy được vận hành trên phần cứng mạnh mẽ hoặc chuyên dụng (như ASIC), và mục tiêu duy nhất của chúng là đào các block mới càng nhanh càng tốt. Thợ đào chỉ có thể làm việc trong các blockchain sử dụng Proof-of-Work (Bằng chứng Công việc), vì đào thực chất là giải các câu đố PoW. Ví dụ, trong các blockchain Proof-of-Stake (Bằng chứng Cổ phần), không có hoạt động đào.

2. **Fullnode**

    Các nút này xác thực các block được khai thác bởi thợ đào và xác minh giao dịch. Để làm được điều này, chúng phải có toàn bộ bản sao của blockchain. Ngoài ra, các nút này còn thực hiện các hoạt động định tuyến, chẳng hạn như giúp các nút khác tìm thấy nhau.

    Mạng lưới cần có nhiều nút đầy đủ vì chính những nút này sẽ đưa ra quyết định: chúng quyết định xem một block hoặc giao dịch có hợp lệ hay không.

3. **SPV (Simplified Payment Verification)**

    SPV là viết tắt của Simplified Payment Verification (Xác minh Thanh toán Đơn giản). Các nút này không lưu trữ toàn bộ bản sao blockchain, nhưng vẫn có thể xác minh các giao dịch (không phải tất cả, mà chỉ một tập hợp con, ví dụ, các giao dịch được gửi đến một địa chỉ cụ thể). Một nút SPV phụ thuộc vào một nút đầy đủ để lấy dữ liệu, và có thể có nhiều nút SPV được kết nối với một nút đầy đủ. SPV giúp các ứng dụng ví trở nên khả thi: người dùng không cần tải xuống toàn bộ blockchain, nhưng vẫn có thể xác minh các giao dịch của mình.

### 2. Đơn giản hóa mạng

Để triển khai mạng lưới trong blockchain, ta cần đơn giản hóa một số thứ. Vấn đề là ta không có nhiều máy tính để mô phỏng một mạng lưới với nhiều nút. Ta sẽ sử dụng Docker để giải quyết vấn đề này, các máy tính (ở đây là các container) sẽ giao tiếp ở cổng 3000.

### 3. Triển khai thực tế

Điều gì sẽ xảy ra khi bạn tải xuống, chẳng hạn như Bitcoin Core và chạy nó lần đầu tiên? Nó phải kết nối với một số nút để tải xuống trạng thái mới nhất của blockchain. Thử nghĩ đến việc máy tính của bạn làm sao biết được một máy tính nào đó khác có lưu trạng thái blockchain để yêu cầu tải xuống? Đó là nút nào?

Việc mã hóa cứng địa chỉ nút trong Bitcoin Core sẽ là một sai lầm: nút có thể bị tấn công hoặc bị tắt, dẫn đến việc các nút mới không thể tham gia mạng. Thay vào đó, trong Bitcoin Core, có các hạt giống DNS được hardcode. Chúng không phải là nút, mà là máy chủ DNS biết địa chỉ của một số nút. Khi bạn khởi động một Bitcoin Core sạch, nó sẽ kết nối với một trong các hạt giống và nhận được danh sách các nút đầy đủ, sau đó nó sẽ tải xuống blockchain từ đó.

Tuy nhiên, trong quá trình triển khai, ta sẽ tập trung hóa. Ta sẽ có ba nút:

1. Nút trung tâm. Đây là nút mà tất cả các nút khác sẽ kết nối tới và là nút sẽ gửi dữ liệu giữa các nút khác. Có thể xem như nút này có vai trò là **Fullnode**
2. **Miner node**. Nút này sẽ lưu trữ các giao dịch mới trong mempool và khi có đủ giao dịch, nó sẽ khai thác một block mới.
3. Một nút ví. Nút này sẽ được sử dụng để gửi coin giữa các ví. Tuy nhiên, không giống như các nút **SPV**, nó sẽ lưu trữ một bản sao đầy đủ của blockchain.

> Ta sẽ hardcode địa chỉ của nút trung tâm để tất cả các nút đều có thể liên lạc với nút trung tâm.

### 4. Quy trình Demo
#### 4.1 Quy ước cách gọi
Trước tiên ta cần quy ước cách gọi để dễ dàng quan sát và hiểu vai trò của từng nút trong demo so với bitcoin:
- **Host**: Là máy tính vật lý có vai trò trung gian để phân phát tài nguyên động thay vì hardcode như trong bitcoin
- **Fullnode**: có vai trò như một nút đầy đủ, và ở trong demo này chúng còn có vai trò là nút trung tâm để tiếp nhận và phân phát dữ liệu
- **Miner**: có vai trò là nút khai thác
- **Wallet**: là các nút ví có vai trò thực hiện giao dịch, nếu bạn dùng tcoin để giao dịch, thiết bị của bạn là một nút Wallet

#### 4.2 Kịch bản
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

#### 4.3 Triển khai

##### 4.3.1 Yêu cầu
- Go & source code (Nếu dựng file nhị phân từ mã nguồn)
- Docker

##### 4.3.2 Chuẩn bị
Trước tiên ta sẽ dùng docker để tạo ra 3 container dựa trên ubuntu image và đặt hostname cho chúng lần lượt là: `fullnode`, `miner` và `wallet`.

Ta sẽ dùng lớp mạng mặc định mà docker gán cho các container (bridge driver). Lần lượt log-on vào các container theo thứ tự sẽ được các địa chỉ như sau:

- fullnode: `172.17.0.2`
- miner: `172.17.0.3`
- wallet: `172.17.0.4`


##### 4.3.3 Thực hiện

###### Fullnode

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

###### Host
Do ta không hardcode genesis block nên sẽ dùng máy vật lý để phân phát genesis block cho các nút. Trước tiên ta sẽ copy genesis block từ container ra máy vật lý:
```bash
docker cp fullnode:/root/blockchain_172.17.0.2.db blockchain_genesis.db
```

###### Wallet
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

###### Fullnode
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

###### Host
Trước khi chuyển sang nút Wallet để gia nhập mạng blockchain, nút Wallet cần biết đâu là blockchain mà nó sẽ gia nhập thông qua genesis block (một lần nữa, genesis block được hardcode!), do đó ta cần copy genesis block từ máy vật lý vào `wallet` container:
```bash
docker cp blockchain_genesis.db wallet:/root/blockchain_172.17.0.4.db
```

###### Wallet
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

###### Host
Được rồi, đã đến lúc các thợ đào gia nhập đội hình! Tương tự như các nút khác, các thợ đào cũng cần phải biết đâu là blockchain chính thống, chúng ta vẫn phải copy genesis block cho các thợ đào:
```bash
docker cp blockchain_genesis.db miner:/root/blockchain_172.17.0.3.db
```

###### Miner
- Trước tiên các thợ đào cần có một địa chỉ ví để nhận phần thưởng khi đào được các block:
    ```bash
    [miner]$ tcoin createwallet
    Your new address: 1BJiAuYqcChkQCMH8LHWWZT177J4PjuHDw
    ```
- **Bắt đầu đào thôi!** dùng flag `-miner` để chỉ định đây là một nút thợ đào trong mạng, các thợ đào sẽ bắt đầu đào một khối khi có từ 2 giao dịch trở lên, hãy khởi động và đợi các nút Wallet giao dịch:
    ```bash
    tcoin startnode -miner 1BJiAuYqcChkQCMH8LHWWZT177J4PjuHDw
    ```

###### Wallet
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
