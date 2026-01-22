curl -X POST "http://localhost:9200/_bulk" \
-H "Content-Type: application/json" \
-d '
{ "index": { "_index": "products", "_id": "1" } }
{ "id":"1","name":"Sony WH-1000XM4","description":"Wireless noise cancelling headphones","brand":"Sony","category":"Audio","price":29999,"rating":4.8,"in_stock":true,"created_at":"2024-01-01" }
{ "index": { "_index": "products", "_id": "2" } }
{ "id":"2","name":"Sony Wired Earphones","description":"High quality wired earphones","brand":"Sony","category":"Audio","price":1999,"rating":4.1,"in_stock":false,"created_at":"2024-01-02" }
{ "index": { "_index": "products", "_id": "3" } }
{ "id":"3","name":"Apple AirPods Pro","description":"Wireless earbuds with ANC","brand":"Apple","category":"Audio","price":24999,"rating":4.7,"in_stock":true,"created_at":"2024-01-03" }
{ "index": { "_index": "products", "_id": "4" } }
{ "id":"4","name":"Apple AirPods 2","description":"Wireless earbuds","brand":"Apple","category":"Audio","price":12999,"rating":4.5,"in_stock":false,"created_at":"2024-01-04" }
{ "index": { "_index": "products", "_id": "5" } }
{ "id":"5","name":"Boat Rockerz 450","description":"Bluetooth on-ear headphones","brand":"Boat","category":"Audio","price":1999,"rating":4.2,"in_stock":true,"created_at":"2024-01-05" }
{ "index": { "_index": "products", "_id": "6" } }
{ "id":"6","name":"Boat Airdopes 131","description":"True wireless earbuds","brand":"Boat","category":"Audio","price":1499,"rating":4.0,"in_stock":false,"created_at":"2024-01-06" }
{ "index": { "_index": "products", "_id": "7" } }
{ "id":"7","name":"Samsung Galaxy Buds Pro","description":"Premium wireless earbuds","brand":"Samsung","category":"Audio","price":15999,"rating":4.6,"in_stock":true,"created_at":"2024-01-07" }
{ "index": { "_index": "products", "_id": "8" } }
{ "id":"8","name":"Samsung Wired Earphones","description":"In-ear wired headphones","brand":"Samsung","category":"Audio","price":999,"rating":3.9,"in_stock":false,"created_at":"2024-01-08" }
{ "index": { "_index": "products", "_id": "9" } }
{ "id":"9","name":"Nike Air Max","description":"Comfortable running shoes","brand":"Nike","category":"Footwear","price":8999,"rating":4.4,"in_stock":true,"created_at":"2024-01-09" }
{ "index": { "_index": "products", "_id": "10" } }
{ "id":"10","name":"Nike Revolution 6","description":"Lightweight running shoes","brand":"Nike","category":"Footwear","price":4999,"rating":4.1,"in_stock":false,"created_at":"2024-01-10" }

{ "index": { "_index": "products", "_id": "11" } }
{ "id":"11","name":"Sony Bluetooth Speaker","description":"Portable wireless speaker","brand":"Sony","category":"Electronics","price":6999,"rating":4.3,"in_stock":true,"created_at":"2024-01-11" }
{ "index": { "_index": "products", "_id": "12" } }
{ "id":"12","name":"Apple Watch SE","description":"Smartwatch with fitness tracking","brand":"Apple","category":"Wearables","price":29999,"rating":4.6,"in_stock":true,"created_at":"2024-01-12" }
{ "index": { "_index": "products", "_id": "13" } }
{ "id":"13","name":"Samsung Galaxy Watch","description":"Smartwatch with AMOLED display","brand":"Samsung","category":"Wearables","price":27999,"rating":4.4,"in_stock":false,"created_at":"2024-01-13" }
{ "index": { "_index": "products", "_id": "14" } }
{ "id":"14","name":"Boat Smart Watch","description":"Affordable smartwatch","brand":"Boat","category":"Wearables","price":2499,"rating":3.8,"in_stock":true,"created_at":"2024-01-14" }
{ "index": { "_index": "products", "_id": "15" } }
{ "id":"15","name":"Nike Training Shoes","description":"Gym and training shoes","brand":"Nike","category":"Footwear","price":5999,"rating":4.2,"in_stock":true,"created_at":"2024-01-15" }

{ "index": { "_index": "products", "_id": "16" } }
{ "id":"16","name":"Sony Neckband","description":"Wireless neckband headphones","brand":"Sony","category":"Audio","price":2999,"rating":4.0,"in_stock":false,"created_at":"2024-01-16" }
{ "index": { "_index": "products", "_id": "17" } }
{ "id":"17","name":"Apple Wired EarPods","description":"Wired earphones with mic","brand":"Apple","category":"Audio","price":1999,"rating":4.1,"in_stock":true,"created_at":"2024-01-17" }
{ "index": { "_index": "products", "_id": "18" } }
{ "id":"18","name":"Samsung Soundbar","description":"Home theatre sound system","brand":"Samsung","category":"Electronics","price":19999,"rating":4.5,"in_stock":true,"created_at":"2024-01-18" }
{ "index": { "_index": "products", "_id": "19" } }
{ "id":"19","name":"Boat Stone Speaker","description":"Portable speaker","brand":"Boat","category":"Electronics","price":3499,"rating":4.0,"in_stock":false,"created_at":"2024-01-19" }
{ "index": { "_index": "products", "_id": "20" } }
{ "id":"20","name":"Nike Slides","description":"Comfortable daily wear slides","brand":"Nike","category":"Footwear","price":1999,"rating":3.9,"in_stock":true,"created_at":"2024-01-20" }
'


curl -X POST "http://localhost:9200/_bulk" \
-H "Content-Type: application/json" \
-d '
{ "index": { "_index": "products", "_id": "21" } }
{ "id":"21","name":"Sony Extra Bass Headphones","description":"Powerful bass wireless headphones","brand":"Sony","category":"Audio","price":4999,"rating":4.3,"in_stock":true,"created_at":"2024-02-01" }
{ "index": { "_index": "products", "_id": "22" } }
{ "id":"22","name":"Sony Walkman Earphones","description":"Lightweight wired earphones","brand":"Sony","category":"Audio","price":1299,"rating":3.9,"in_stock":false,"created_at":"2024-02-02" }

{ "index": { "_index": "products", "_id": "23" } }
{ "id":"23","name":"Apple HomePod Mini","description":"Smart speaker with Siri","brand":"Apple","category":"Electronics","price":9999,"rating":4.4,"in_stock":true,"created_at":"2024-02-03" }
{ "index": { "_index": "products", "_id": "24" } }
{ "id":"24","name":"Apple MagSafe Charger","description":"Wireless magnetic charger","brand":"Apple","category":"Electronics","price":4499,"rating":4.2,"in_stock":false,"created_at":"2024-02-04" }

{ "index": { "_index": "products", "_id": "25" } }
{ "id":"25","name":"Boat BassHeads 100","description":"Wired in-ear headphones","brand":"Boat","category":"Audio","price":399,"rating":3.8,"in_stock":true,"created_at":"2024-02-05" }
{ "index": { "_index": "products", "_id": "26" } }
{ "id":"26","name":"Boat Party Speaker","description":"Loud party bluetooth speaker","brand":"Boat","category":"Electronics","price":7999,"rating":4.1,"in_stock":false,"created_at":"2024-02-06" }

{ "index": { "_index": "products", "_id": "27" } }
{ "id":"27","name":"Samsung Galaxy S23","description":"Flagship Android smartphone","brand":"Samsung","category":"Electronics","price":74999,"rating":4.7,"in_stock":true,"created_at":"2024-02-07" }
{ "index": { "_index": "products", "_id": "28" } }
{ "id":"28","name":"Samsung Charger","description":"Fast charging adapter","brand":"Samsung","category":"Electronics","price":1999,"rating":4.0,"in_stock":false,"created_at":"2024-02-08" }

{ "index": { "_index": "products", "_id": "29" } }
{ "id":"29","name":"Nike Running Shorts","description":"Breathable sports shorts","brand":"Nike","category":"Footwear","price":2499,"rating":4.0,"in_stock":true,"created_at":"2024-02-09" }
{ "index": { "_index": "products", "_id": "30" } }
{ "id":"30","name":"Nike Sports Socks","description":"Comfortable ankle socks","brand":"Nike","category":"Footwear","price":799,"rating":3.7,"in_stock":false,"created_at":"2024-02-10" }

{ "index": { "_index": "products", "_id": "31" } }
{ "id":"31","name":"Sony Smart TV","description":"55 inch 4K Ultra HD TV","brand":"Sony","category":"Electronics","price":64999,"rating":4.6,"in_stock":true,"created_at":"2024-02-11" }
{ "index": { "_index": "products", "_id": "32" } }
{ "id":"32","name":"Sony TV Remote","description":"Replacement remote control","brand":"Sony","category":"Electronics","price":1499,"rating":3.8,"in_stock":false,"created_at":"2024-02-12" }

{ "index": { "_index": "products", "_id": "33" } }
{ "id":"33","name":"Apple iPhone 14","description":"iOS smartphone with A15 chip","brand":"Apple","category":"Electronics","price":69999,"rating":4.8,"in_stock":true,"created_at":"2024-02-13" }
{ "index": { "_index": "products", "_id": "34" } }
{ "id":"34","name":"Apple Lightning Cable","description":"USB to lightning charging cable","brand":"Apple","category":"Electronics","price":1999,"rating":4.1,"in_stock":false,"created_at":"2024-02-14" }

{ "index": { "_index": "products", "_id": "35" } }
{ "id":"35","name":"Boat Fitness Band","description":"Fitness tracker with heart rate","brand":"Boat","category":"Wearables","price":2999,"rating":3.9,"in_stock":true,"created_at":"2024-02-15" }
{ "index": { "_index": "products", "_id": "36" } }
{ "id":"36","name":"Boat Charging Dock","description":"Smartwatch charging dock","brand":"Boat","category":"Wearables","price":999,"rating":3.5,"in_stock":false,"created_at":"2024-02-16" }

{ "index": { "_index": "products", "_id": "37" } }
{ "id":"37","name":"Samsung Fitness Band","description":"AMOLED fitness tracker","brand":"Samsung","category":"Wearables","price":3999,"rating":4.2,"in_stock":true,"created_at":"2024-02-17" }
{ "index": { "_index": "products", "_id": "38" } }
{ "id":"38","name":"Samsung Watch Strap","description":"Silicone replacement strap","brand":"Samsung","category":"Wearables","price":1299,"rating":3.6,"in_stock":false,"created_at":"2024-02-18" }

{ "index": { "_index": "products", "_id": "39" } }
{ "id":"39","name":"Nike Air Zoom Pegasus","description":"Professional running shoes","brand":"Nike","category":"Footwear","price":11999,"rating":4.5,"in_stock":true,"created_at":"2024-02-19" }
{ "index": { "_index": "products", "_id": "40" } }
{ "id":"40","name":"Nike Shoe Cleaner","description":"Footwear cleaning kit","brand":"Nike","category":"Footwear","price":999,"rating":3.8,"in_stock":false,"created_at":"2024-02-20" }

{ "index": { "_index": "products", "_id": "41" } }
{ "id":"41","name":"Sony Gaming Headset","description":"Surround sound gaming headset","brand":"Sony","category":"Audio","price":8999,"rating":4.4,"in_stock":true,"created_at":"2024-02-21" }
{ "index": { "_index": "products", "_id": "42" } }
{ "id":"42","name":"Sony AUX Cable","description":"3.5mm audio cable","brand":"Sony","category":"Audio","price":499,"rating":3.6,"in_stock":false,"created_at":"2024-02-22" }

{ "index": { "_index": "products", "_id": "43" } }
{ "id":"43","name":"Apple Studio Headphones","description":"Over-ear professional headphones","brand":"Apple","category":"Audio","price":39999,"rating":4.7,"in_stock":true,"created_at":"2024-02-23" }
{ "index": { "_index": "products", "_id": "44" } }
{ "id":"44","name":"Apple Ear Cushions","description":"Replacement ear cushions","brand":"Apple","category":"Audio","price":5999,"rating":4.0,"in_stock":false,"created_at":"2024-02-24" }

{ "index": { "_index": "products", "_id": "45" } }
{ "id":"45","name":"Boat Immortal Gaming Headset","description":"RGB gaming headphones","brand":"Boat","category":"Audio","price":3499,"rating":4.1,"in_stock":true,"created_at":"2024-02-25" }
{ "index": { "_index": "products", "_id": "46" } }
{ "id":"46","name":"Boat Mic Extension","description":"Microphone extension cable","brand":"Boat","category":"Audio","price":699,"rating":3.4,"in_stock":false,"created_at":"2024-02-26" }

{ "index": { "_index": "products", "_id": "47" } }
{ "id":"47","name":"Samsung Monitor 27 inch","description":"Full HD IPS monitor","brand":"Samsung","category":"Electronics","price":14999,"rating":4.3,"in_stock":true,"created_at":"2024-02-27" }
{ "index": { "_index": "products", "_id": "48" } }
{ "id":"48","name":"Samsung HDMI Cable","description":"High speed HDMI cable","brand":"Samsung","category":"Electronics","price":899,"rating":3.7,"in_stock":false,"created_at":"2024-02-28" }

{ "index": { "_index": "products", "_id": "49" } }
{ "id":"49","name":"Nike Gym Bag","description":"Durable sports gym bag","brand":"Nike","category":"Footwear","price":2999,"rating":4.0,"in_stock":true,"created_at":"2024-03-01" }
{ "index": { "_index": "products", "_id": "50" } }
{ "id":"50","name":"Nike Shoe Laces","description":"Replacement shoe laces","brand":"Nike","category":"Footwear","price":399,"rating":3.5,"in_stock":false,"created_at":"2024-03-02" }
'
