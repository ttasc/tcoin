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

> ##### Có thể bạn chưa biết:
> Khối genesis của bitcoin luôn được **mã hóa cứng** (hardcode) vào các phần mềm sử dụng chuỗi khối của bitcoin.

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

`ca07ca` là giá trị thập lục phân của bộ đếm, tức là 13240266 trong hệ thập phân.


## Phần tiếp theo đang cập nhật...
---
## Demo

![demo](./demo/demo.md)
