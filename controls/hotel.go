package controls

import (
	"fmt"

	"strconv"

	"checkin/config"
	"checkin/models"

	"github.com/gin-gonic/gin"
)

//Admin adding the hotel
func AddHotel(c *gin.Context) {
	var addhotel models.Hotel
	if c.Bind(&addhotel) != nil {
		c.JSON(400, gin.H{
			"Error": "Could not bind JSON data",
		})
		return
	}
	fmt.Println("================", addhotel.HotelName)
	DB := config.DB
	result := DB.Create(&addhotel)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Message":       "New Hotel added Successfully",
		"Hotel details": addhotel,
	})
}

//view hotel
func ViewHotel(c *gin.Context) {
	var hotelData []models.Hotel
	db := config.DB
	result := db.First(&hotelData)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Message": "Hotel is empty",
		})
		return
	}
	c.JSON(200, gin.H{
		"Hotels data": hotelData,
	})
}

//Edit hotel by admin
func EditHotel(c *gin.Context) {
	bid := c.Param("id")
	id, err := strconv.Atoi(bid)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error in string conversion",
		})
	}
	var edithotels models.Hotel
	if c.Bind(&edithotels) != nil {
		c.JSON(400, gin.H{
			"Error": "Error in binding the JSON data",
		})
		return
	}
	edithotels.ID = uint(id)
	DB := config.DB

	result := DB.Model(&edithotels).Updates(models.Hotel{
		HotelName: edithotels.HotelName,
		Location:  edithotels.Location,
		Phone:    edithotels.Phone,
	})

	if result.Error != nil {
		c.JSON(404, gin.H{
			"Error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Successfully updated the Hotel",
	})
}

// View Hotels and the rooms
// func ViewHotelRooms(c *gin.Context) {
//   limit, _ := strconv.Atoi(c.Query("limit"))
//   offset, _ := strconv.Atoi(c.Query("offset"))

//   type Room struct {
//     RoomID     uint   `json:"Room_id"`
//     RoomName   string `json:"room_name"`
//     Facilities string `json:"facilities"`
//     Stock      uint   `json:"stock"`
//     Available  uint   `json:"available"`
//     Price      uint   `json:"price"`
//     HotelID    uint   `json:"Hotel_id"` // Renamed for consistency with model
//   }

//   type Hotel struct {
//     ID        uint   `json:"id"`
//     HotelName string `json:"hotel_name"`
//     Location  string `json:"location"`
//     Phone     string `json:"phone"`
//     Rooms     []Room `json:"rooms"`
//   }

//   var hotels []Hotel

//   db := config.DB
//   query := `SELECT rooms.room_id, rooms.room_name, rooms.facilities, rooms.stock, rooms.avaliable, rooms.price, hotels.id, hotels.hotel_name, hotels.location, hotels.phone
//             FROM rooms
//             LEFT JOIN hotels ON rooms.hotel_id=hotels.id
//             GROUP BY hotels.id, rooms.room_id` // Group by hotel.id first

//   if limit != 0 || offset != 0 {
//     if limit == 0 {
//       query = fmt.Sprintf("%s OFFSET %d", query, offset)
//     } else if offset == 0 {
//       query = fmt.Sprintf("%s LIMIT %d", query, limit)
//     } else {
//       query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
//     }
//   }

//   rows, err := db.Raw(query).Rows()
//   if err != nil {
//     c.JSON(404, gin.H{
//       "Error": err.Error(),
//     })
//     return
//   }
//   defer rows.Close()

//   var currentHotelID uint
//   var currentHotelData Hotel

//   for rows.Next() {
//     var room Room
//     var hotel Hotel
//     err := rows.Scan(&room.RoomID, &room.RoomName, &room.Facilities, &room.Stock, &room.Available, &room.Price, &hotel.ID, &hotel.HotelName, &hotel.Location, &hotel.Phone)
//     if err != nil {
//       c.JSON(404, gin.H{
//         "Error": err.Error(),
//       })
//       return
//     }

//     if currentHotelID != currentHotelData.ID {
//       // New hotel encountered, start a new record
//       currentHotelData = Hotel{
//         ID:        hotel.ID,
//         HotelName: hotel.HotelName,
//         Location:  hotel.Location,
//         Phone:     hotel.Phone,
//         Rooms:     []Room{},
//       }
//       hotels = append(hotels, currentHotelData)
//     }

//     currentHotelData.Rooms = append(currentHotelData.Rooms, room)
//     currentHotelID = hotel.ID
//   }

//   c.JSON(200, gin.H{
//     "hotels": hotels,
//   })
// }

// 	//fetching the user id from the tocken
// 	id, err := strconv.Atoi(c.GetString("userid"))
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Error": "Error in string conversion",
// 		})
// 		return
// 	}

// 	db := config.DB

// 	//checking the product is exist or not
// 	result := db.First(&productData, bindData.Product_id)
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Message": "Product not exist",
// 		})
// 		return
// 	}

// 	//checking stock quantity
// 	if bindData.Quantity > productData.Stock {
// 		c.JSON(404, gin.H{
// 			"Message": "Out of Stock",
// 		})
// 		return
// 	}

// 	var sum uint
// 	var Price uint

// 	//checking the produt_id and user_id  in the carts table
// 	err = db.Table("carts").Where("product_id = ? AND userid = ? ", bindData.Product_id, id).Select("quantity", "total_price").Row().Scan(&sum, &Price)
// 	fmt.Println("this is the erro : ", err)
// 	if err != nil {
// 		totalprice := productData.Price * bindData.Quantity
// 		cartitems := models.Cart{
// 			ProductId:  bindData.Product_id,
// 			Quantity:   bindData.Quantity,
// 			Price:      productData.Price,
// 			TotalPrice: totalprice,
// 			Userid:     uint(id),
// 		}

// 		//Creating the table carts
// 		result := db.Create(&cartitems)
// 		if result.Error != nil {
// 			c.JSON(400, gin.H{
// 				"Error": result.Error.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(200, gin.H{
// 			"Message": "Added to the Cart Successfull",
// 		})
// 		return
// 	}

// 	//calculatin the tottal quantity and total Price
// 	totalQuantity := sum + bindData.Quantity
// 	totalPrice := productData.Price * totalQuantity

// 	//updating the quatity and the total price  to the carts
// 	result = db.Model(&models.Cart{}).Where("product_id = ? AND userid = ? ", bindData.Product_id, id).Updates(map[string]interface{}{"quantity": totalQuantity, "total_price": totalPrice})
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"Message": "Quantity added Successfully",
// 	})
// }

// //View cart items using user id
// func ViewCart(c *gin.Context) {
// 	id, err := strconv.Atoi(c.GetString("userid"))
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Error": "Error in string conversion",
// 		})
// 	}

// 	type cartdata struct {
// 		Product_name string
// 		Quantity     uint
// 		Total_price  uint
// 		Image        string
// 		Price        string
// 	}

// 	var datas []cartdata

// 	db := config.DB

// 	result := db.Table("carts").
// 		Select("products.product_name, images.image, carts.quantity, carts.price, carts.total_price").
// 		Joins("INNER JOIN products ON products.product_id=carts.product_id").
// 		Joins("INNER JOIN images ON images.product_id=carts.product_id").
// 		Where("userid = ?", id).
// 		Scan(&datas)

// 	if result.Error != nil {
// 		c.JSON(404, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}
// 	fmt.Println("this is carts data : ", datas)
// 	if datas != nil {
// 		c.JSON(200, gin.H{
// 			"Cart Items": datas,
// 		})
// 	} else {
// 		c.JSON(404, gin.H{
// 			"Message": "Cart is empty",
// 		})
// 	}

// }

// //Delete cart of a perticular user id
// func DeleteCart(c *gin.Context) {
// 	id := c.Param("id")
// 	userid, err := strconv.Atoi(c.GetString("userid"))
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Error": "Error in string conversion",
// 		})
// 	}

// 	db := config.DB
// 	result := db.Exec("delete from carts where id= ? AND userid = ?", id, userid)
// 	count := result.RowsAffected
// 	if count == 0 {
// 		c.JSON(400, gin.H{
// 			"Message": "Cart not exist",
// 		})
// 		return
// 	}
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"Cart Items": "Delete successfully",
// 	})
// }

//Add image of the product by admin
// func AddImages(c *gin.Context) {
// 	imagepath, _ := c.FormFile("image")
// 	pid, _ := strconv.Atoi(c.PostForm("product_id"))

// 	db := config.DB
// 	var product models.Product
// 	result := db.First(&product, pid)
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	extension := filepath.Ext(imagepath.Filename)
// 	image := uuid.New().String() + extension
// 	c.SaveUploadedFile(imagepath, "./public/images"+image)

// 	imagedata := models.Image{
// 		Image:     image,
// 		ProductId: uint(pid),
// 	}
// 	result = db.Create(&imagedata)
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"Message": "Image Added Successfully",
// 	})
// }

//Searching the product using product name and Hotel name. If product name does't exist then it search using the Hotel name
// func SearchProduct(c *gin.Context) {
// 	type Data struct {
// 		SearchValue string
// 	}
// 	type product struct {
// 		Product_name string
// 		Description  string
// 		Price        float64
// 	}

// 	var userEnterData Data
// 	if c.Bind(&userEnterData) != nil {
// 		c.JSON(400, gin.H{
// 			"Error": "countl not bind the JSON data",
// 		})
// 	}
// 	var products []product
// 	db := config.DB
// 	var count int64
// 	result := db.Raw("SELECT product_name,description,price FROM rooms WHERE Hotel_id IN (SELECT id FROM Hotels WHERE Hotel_name LIKE ?)", "%"+userEnterData.SearchValue+"%").Scan(&products).Count(&count)
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": "Product not exist",
// 		})
// 		return
// 	}

// 	if count <= 0 {
// 		result := db.Raw("SELECT * FROM products WHERE product_name LIKE ?", "%"+userEnterData.SearchValue+"%").Find(&products)
// 		if result.Error != nil {
// 			c.JSON(400, gin.H{
// 				"Error": result.Error.Error(),
// 			})
// 		}
// 		return
// 	}

// 	if count == 0 {
// 		c.JSON(400, gin.H{
// 			"Message": "Product not exist",
// 		})
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"products": products,
// 	})
// }

// creating pdf file containing invoice for show to the user
// type Invoice struct {
// 	Name          string
// 	Email         string
// 	PaymentMethod string
// 	Totalamount   int64
// 	Date          string
// 	OrderId       uint
// 	Address       []Address
// 	Items         []Item
// }
// type Address struct {
// 	Phoneno  string
// 	Houseno  string
// 	Area     string
// 	Landmark string
// 	City     string
// 	Pincode  string
// 	District string
// 	State    string
// 	Country  string
// }

// type Item struct {
// 	Product     string
// 	Description string
// 	Qty         uint
// 	Price       uint
// }

// //Templates for creating the pdf
// const invoiceTemplate = `
// Order ID : {{.OrderId}}<br>
// Order Date:{{.Date}} <br><hr>
// Name : {{.Name}} <br>
// Email: {{.Email}}<br>
// <hr>
// Billing Address :
// {{range .Address}}

// Phone number : {{.Phoneno}} <br>
// House number : {{.Houseno}} <br>
// Area : {{.Area}} <br>
// Landmark : {{.Landmark}} <br>
// City : {{.City}} <br>
// Pincode : {{.Pincode}} <br>
// District : {{.District}} <br>
// State : {{.State}} <br>
// Country : {{.Country}} <br>
// {{end}}
// <hr>
// Payment method : {{.PaymentMethod}}<br>
// <hr>
// {{range .Items}}

// Product :{{.Product}}  <br>
// Description: {{.Description}}<br>
// Price : {{.Price}}<br><br>

// {{end}}
// <hr><br>
// Total Amount : {{.Totalamount}}<br>
// `

// func InvoiceF(c *gin.Context) {
// 	fmt.Println()

// 	id, err := strconv.Atoi(c.GetString("userid"))
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Error": "Error in string conversion",
// 		})
// 		return
// 	}

// 	db := config.DB
// 	var user models.User
// 	var Payment models.Payment
// 	var oderData models.OderDetails
// 	var address models.Address
// 	var Oder_item models.Oder_item

// 	//fetching the data from table Oder_item using usder id
// 	result := db.Last(&Oder_item).Where("useridno = ?", id)
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	//fetching the data from table Oder_details using userid and oder_idtemid, for fetching the oder_itemid
// 	result = db.Last(&oderData).Where("useridno = ? AND oder_itemid = ?", id, Oder_item.OrderId)
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	//Fetching the data from table users using userid
// 	result = db.First(&user, id)
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	//fetching the user address using address id from table Oder_Details
// 	result = db.First(&address, oderData.AddressId)
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	//fetching the payment detail form table Payments using userid
// 	result = db.Last(&Payment, "user_id = ?", id)
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	//fetching the product data from table products using Oder_itemid from table Oder_item.
// 	var products []models.Product
// 	err = db.Joins("JOIN oder_details ON products.product_id = oder_details.product_id").
// 		Where("oder_details.oder_item_id = ?", oderData.OderItemId).Find(&products).Error
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Error": "somthing went wrong",
// 		})
// 		return
// 	}

// 	//To list the product details from products, product data assign to slice items
// 	items := make([]Item, len(products))
// 	for i, data := range products {
// 		items[i] = Item{
// 			Product:     data.ProductName,
// 			Price:       data.Price,
// 			Description: data.Description,
// 		}
// 	}

// 	//spliting the date from time.time
// 	timeString := Payment.Date.Format("2006-01-02")

// 	//executing the template Invoice
// 	invoice := Invoice{
// 		Name:          user.FirstName,
// 		Date:          timeString,
// 		Email:         user.Email,
// 		OrderId:       oderData.OderItemId,
// 		PaymentMethod: Payment.PaymentMethod,
// 		Totalamount:   int64(Payment.Totalamount),
// 		Address: []Address{
// 			{
// 				Phoneno:  address.Phoneno,
// 				Houseno:  address.Houseno,
// 				Area:     address.Area,
// 				Landmark: address.Landmark,
// 				City:     address.City,
// 				Pincode:  address.Pincode,
// 				District: address.District,
// 				State:    address.State,
// 				Country:  address.Country,
// 			},
// 		},
// 		Items: items,
// 	}

// 	tmpl, err := template.New("invoice").Parse(invoiceTemplate)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Error": err.Error(),
// 		})
// 		return
// 	}

// 	var buf bytes.Buffer
// 	err = tmpl.Execute(&buf, invoice)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Error": err.Error(),
// 		})
// 		return
// 	}

// 	cmd := exec.Command("wkhtmltopdf", "-", "./public/invoice.pdf")
// 	cmd.Stdin = &buf
// 	err = cmd.Run()
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Error": err.Error(),
// 		})
// 		return
// 	}

// 	c.HTML(200, "invoice.html", gin.H{})
// }

// //To download the pdf file
// func Download(c *gin.Context) {
// 	c.Header("Content-Disposition", "attachment; filename=invoice.pdf")
// 	c.Header("Content-Type", "application/pdf")
// 	c.File("invoice.pdf")
// }

// //<<<<<<<<<<<<< Sales Report >>>>>>>>>>>>>>>>>>>>>>>>
// func SalesReport(c *gin.Context) {

// 	//fetching the dates from the URL
// 	startDate := c.Query("startDate")
// 	endDateStr := c.Query("endDate")

// 	//converting the dates string to time.time
// 	fromTime, err := time.Parse("2006-01-02", startDate)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "Invalid start Date",
// 		})
// 		return
// 	}
// 	toTime, err := time.Parse("2006-01-02", endDateStr)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "Invalid end Date",
// 		})
// 		return
// 	}

// 	//fetching the data from the table Order details where start date to end date
// 	var orderDetail []models.OderDetails
// 	// var reportData []Report
// 	db := config.DB

// 	result := db.Preload("Product").Preload("Payment").
// 		Where("created_at BETWEEN ? AND ?", fromTime, toTime).
// 		Find(&orderDetail)

// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	f := excelize.NewFile()

// 	// Create a new sheet.
// 	SheetName := "Sheet1"
// 	index := f.NewSheet(SheetName)

// 	// Set the value of headers
// 	f.SetCellValue(SheetName, "A1", "Order Date")
// 	f.SetCellValue(SheetName, "B1", "Order ID")
// 	f.SetCellValue(SheetName, "C1", "Product name")
// 	f.SetCellValue(SheetName, "D1", "Price")
// 	f.SetCellValue(SheetName, "E1", "Total Amount")
// 	f.SetCellValue(SheetName, "F1", "Payment method")
// 	f.SetCellValue(SheetName, "G1", "Payment Status")
// 	// Set the value of cell
// 	for i, report := range orderDetail {
// 		row := i + 2
// 		f.SetCellValue(SheetName, fmt.Sprintf("A%d", row), report.CreatedAt.Format("01/02/2006"))
// 		f.SetCellValue(SheetName, fmt.Sprintf("B%d", row), report.Oderid)
// 		f.SetCellValue(SheetName, fmt.Sprintf("C%d", row), report.Product.ProductName)
// 		f.SetCellValue(SheetName, fmt.Sprintf("D%d", row), report.Product.Price)
// 		f.SetCellValue(SheetName, fmt.Sprintf("E%d", row), report.Payment.Totalamount)
// 		f.SetCellValue(SheetName, fmt.Sprintf("F%d", row), report.Payment.PaymentMethod)
// 		f.SetCellValue(SheetName, fmt.Sprintf("G%d", row), report.Payment.Status)

// 	}

// 	// Set active sheet of the workbook.
// 	f.SetActiveSheet(index)

// 	// Save the Excel file with the name "test.xlsx".
// 	if err := f.SaveAs("./public/SalesReport.xlsx"); err != nil {
// 		fmt.Println(err)
// 	}
// 	CovertingExelToPdf(c)
// 	c.HTML(200, "SalseReport.html", gin.H{})

// }

// func CovertingExelToPdf(c *gin.Context) {
// 	// Open the Excel file
// 	xlFile, err := xlsx.OpenFile("./public/SalesReport.xlsx")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Create a new PDF document
// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()
// 	pdf.SetFont("Arial", "", 14)

// 	// Convertig each cell in the Excel file to a PDF cell
// 	for _, sheet := range xlFile.Sheets {
// 		for _, row := range sheet.Rows {
// 			for _, cell := range row.Cells {
// 				//if there is any empty cell values skiping that
// 				if cell.Value == "" {
// 					continue
// 				}

// 				pdf.Cell(40, 10, cell.Value)
// 			}
// 			pdf.Ln(-1)
// 		}
// 	}

// 	// Save the PDF document
// 	err = pdf.OutputFileAndClose("./public/SalesReport.pdf")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// }

// func DownloadExel(c *gin.Context) {
// 	c.Header("Content-Disposition", "attachment; filename=SalesReport.xlsx")
// 	c.Header("Content-Type", "application/xlsx")
// 	c.File("./public/SalesReport.xlsx")
// }

// func Downloadpdf(c *gin.Context) {
// 	c.Header("Content-Disposition", "attachment; filename=SalesReport.pdf")
// 	c.Header("Content-Type", "application/pdf")
// 	c.File("./public/SalesReport.pdf")
// }

// //Wallet history
// func WalletHistory(c *gin.Context) {
// 	userid, err := strconv.Atoi(c.GetString("userid"))
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Error": err.Error(),
// 		})
// 	}
// 	var WalletHistory []models.WalletHistory
// 	db := config.DB
// 	result := db.Find(&WalletHistory).Where("user_id", userid)
// 	if result.Error != nil {
// 		c.JSON(400, gin.H{
// 			"Error": result.Error.Error(),
// 		})
// 		return
// 	}
// 	var response []map[string]interface{}
// 	for _, history := range WalletHistory {
// 		row := map[string]interface{}{
// 			"Amount":          history.Amount,
// 			"TransactionType": history.TransctionType,
// 			"Date":            history.Date,
// 		}
// 		response = append(response, row)
// 	}

// 	c.JSON(200, gin.H{"data": response})
// }
