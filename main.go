package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Product представляет продукт
type Product struct {
	ID             int
	Title          string
	ImageURL       string
	Name           string
	Price          float64
	Description    string
	Specifications string
	Quantity       int
}

// Пример списка продуктов
var products = []Product{
	{ID: 1, Title: "RTX 4060Ti Windforce OC 8G", ImageURL: "https://cdn1.ozone.ru/s3/multimedia-e/6660780194.jpg", Name: "RTX 4060Ti Windforce OC 8G", Price: 41540, Description: "Видеокарта Gigabyte RTX4060 WINDFORCE OC 8GB GDDR6 128-bit DPx2 HDMIx2 2FAN RTL Прогрессивная микроархитектура Ada Lovelace, фирменная технология NVIDIA DLSS 3 и полноценная реализация трасировки лучей Тензорные ядра 4-поколения: прирост производительности с DLSS 3 до 4x (по сравнению с типовой процедурой рендеринга сцены) RT-ядра 3-поколения: 2-кратный прирост производительности на операциях трассировки лучей Графический процессор GeForce RTX 4060 ВидеоОЗУ GDDR6 8 Гбайт, 128-разрядная шина памяти Система охлаждения WINDFORCE Защитная пластина на тыльной стороне печатной платы.", Specifications: "1", Quantity: 0},
	{ID: 2, Title: "RTX 4070 EAGLE OC 12G", ImageURL: "https://cdn1.ozone.ru/s3/multimedia-1-f/7036023075.jpg", Name: "RTX 4070 EAGLE OC 12G", Price: 84480, Description: "Видеокарта GIGABYTE GeForce RTX 4070 Ti EAGLE OC 12GB - это продукт от известного производителя GIGABYTE, который зарекомендовал себя на рынке компьютерной техники. Видеокарта оснащена видеопроцессором NVIDIA GeForce RTX 4070 Ti, который обеспечивает высокую производительность и реалистичное изображение. Техпроцесс составляет 4 нм, что гарантирует высокую скорость обработки данных и низкое энергопотребление. Объем видеопамяти составляет 12 ГБ, что позволяет работать с тяжелыми графическими приложениями и играми. Тип памяти GDDR6X обеспечивает высокую скорость передачи данных и стабильность работы.", Specifications: "1", Quantity: 0},
	{ID: 3, Title: "RTX 4090 Windforce 24G", ImageURL: "https://avatars.mds.yandex.net/get-mpic/10229228/2a00000190e682a74ff9549f8bf50a7612b6/orig", Name: "RTX 4090 Windforce 24G", Price: 304920, Description: "Видеокарта GIGABYTE GeForce RTX 4090 WINDFORCE V2 [GV-N4090WF3V2-24GD] на основе архитектуры NVIDIA Ada Lovelace обеспечивает высокую графическую производительность для работы с программами и запуска игр на ПК. Процессор функционирует с частотой 2230 МГц, которая способна повышаться до значения 2520 МГц в режиме разгона. Видеокарта оснащена 24 ГБ памяти стандарта GDDR6X с пропускной способностью 1008 Гбайт/сек, что обеспечивает быстродействие обработки графических данных.", Specifications: "1", Quantity: 0},
	{ID: 4, Title: "Видеокарта Gigabyte Radeon RX 7600 GAMING OC 8G", ImageURL: "https://avatars.mds.yandex.net/get-mpic/11225627/2a0000018af87c2c8e082045cc24dd2cdac3/180x240", Name: "Видеокарта Gigabyte Radeon RX 7600 GAMING OC 8G", Price: 33227, Description: "Бренд: GIGABYTE Тип поставки: Ret Количество вентиляторов: 3 Цвет: черный Для геймеров: ДА PartNumber/Артикул Производителя: GV-R76GAMING OC-8GD Тип: Видеокарта Длина упаковки (ед): 0.41", Specifications: "Бренд: GIGABYTE Тип поставки: Ret Количество вентиляторов: 3 Цвет: черный Для геймеров: ДА PartNumber/Артикул Производителя: GV-R76GAMING OC-8GD Тип: Видеокарта Длина упаковки (ед): 0.41", Quantity: 0},
	{ID: 5, Title: "Видеокарта Acer RX7700XT NITRO OC 12GB GDDR6 192bit 3xDP HDMI 2FAN RTL", ImageURL: "https://avatars.mds.yandex.net/get-mpic/5245452/2a00000192b631bf90280871d2e4c4d9695b/74x100", Name: "Видеокарта Acer RX7700XT NITRO OC 12GB GDDR6 192bit 3xDP HDMI 2FAN RTL", Price: 47767, Description: "Эта видеокарта может обеспечить высокую производительность в современных играх и других графических приложениях. Она имеет достаточный объём видеопамяти и широкий интерфейс для подключения нескольких мониторов одновременно. Охлаждение с помощью двух вентиляторов обеспечивает стабильную работу при высоких нагрузках.", Specifications: "* **Графический процессор:** AMD Radeon RX 7700 XT. * **Объём видеопамяти:** 12 ГБ GDDR6. * **Ширина шины памяти:** 192 бита.", Quantity: 0},
}

// обработчик для GET-запроса, возвращает список продуктов
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	json.NewEncoder(w).Encode(products)
}

// обработчик для POST-запроса, добавляет продукт
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received new Product: %+v\n", newProduct)
	var lastID int = len(products)

	for _, productItem := range products {
		if productItem.ID > lastID {
			lastID = productItem.ID
		}
	}
	newProduct.ID = lastID + 1
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

//Добавление маршрута для получения одного продукта

func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем продукт с данным ID
	for _, Product := range products {
		if Product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Product)
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// удаление продукта по id
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем и удаляем продукт с данным ID
	for i, Product := range products {
		if Product.ID == id {
			// Удаляем продукт из среза
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // Успешное удаление, нет содержимого
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Обновление продукта по id
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Декодируем обновлённые данные продукта
	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ищем продукт для обновления
	for i, Product := range products {
		if Product.ID == id {

			products[i].Title = updatedProduct.Title
			products[i].ImageURL = updatedProduct.ImageURL
			products[i].Name = updatedProduct.Name
			products[i].Price = updatedProduct.Price
			products[i].Description = updatedProduct.Description
			products[i].Specifications = updatedProduct.Specifications
			products[i].Quantity = updatedProduct.Quantity

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/products", getProductsHandler)           // Получить все продукты
	http.HandleFunc("/products/create", createProductHandler)  // Создать продукт
	http.HandleFunc("/products/", getProductByIDHandler)       // Получить продукт по ID
	http.HandleFunc("/products/update/", updateProductHandler) // Обновить продукт
	http.HandleFunc("/products/delete/", deleteProductHandler) // Удалить продукт

	fmt.Println("Server is running on http://localhost:8080 !")
	http.ListenAndServe(":8080", nil)
}
