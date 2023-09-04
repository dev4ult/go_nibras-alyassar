package main

type kendaraan struct {
	totalRoda int

	kecepatanPerJam int
}

type mobil struct {
	kendaraan
}

func (objekMobil *mobil) berjalan() {

	objekMobil.tambahKecepatan(10)

}

func (objekMobil *mobil) tambahKecepatan(kecepatanBaru int) {

	objekMobil.kecepatanPerJam = objekMobil.kecepatanPerJam + kecepatanBaru

}

func main() {

	mobilCepat := mobil{}

	mobilCepat.berjalan()

	mobilCepat.berjalan()

	mobilCepat.berjalan()

	mobilLamban := mobil{}

	mobilLamban.berjalan()

}