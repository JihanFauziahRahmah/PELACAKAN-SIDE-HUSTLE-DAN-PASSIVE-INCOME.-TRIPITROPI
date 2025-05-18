package main

import "fmt"

// Struct untuk menyimpan data pendapatan
type Pendapatan struct {
	Sumber   string
	Jumlah   float64
	Tanggal  string
	Kategori string
}

type ID struct {
	Username       string
	PassUser       string
	Kategori       string
	Target         int
	dataPendapatan arrPendapatan
}
type detail struct {
	nama, deskripsi       string
	jumlah                int
	tanggal, bulan, tahun int
}

const NMAX int = 50

type arrID [NMAX]ID
type arrPendapatan [NMAX]detail

func main() {
	var pilih, nArrayPengguna int
	var arrayPengguna arrID

	for {
		menu()
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			menuRegister(&arrayPengguna, &nArrayPengguna)
		case 2:
			menuLogin(&arrayPengguna, &nArrayPengguna)
		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

//fungsi untuk tampilan login
func tampilkanDashboard(user *ID) {
	fmt.Println("\n=== Aplikasi Pencatatan Pendapatan ===")
	fmt.Println("1. Tambah Pendapatan Baru")
	fmt.Println("2. Tampilkan Laporan Keuangan")
	fmt.Println("3. Tampilkan Analisis Terkini")
	fmt.Println("4. Tampilkan Progres Target")
	fmt.Println("5. Logout")
	fmt.Println("Pilih menu: (1/2/3/4/5): ")

	var pilihan int
	n := hitungJumlahPendapatan(user)

	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tambahPendapatan(user, &n)
	case 2:
		laporan(user, &n)
	case 3:
		TampilanAnalisis()
	case 4:
		TampilanProgres(user, &n)
	case 5:
		return
	}
}

func hitungJumlahPendapatan(user *ID) int {
	// hitung jumlah data pendapatan yang sudah diisi (jumlah != 0)
	count := 0
	for i := 0; i < NMAX; i++ {
		if user.dataPendapatan[i].jumlah != 0 {
			count++
		}
	}
	return count
}

//fungsi untuk tampilan utama menu
func menu() {
	fmt.Println("========== Selamat Datang =========")
	fmt.Println("======= Di Aplikasi Pelacak =======")
	fmt.Println("== Side Hustle dan Pssive Income ==")
	fmt.Println("Daftar Menu")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Printf("%s ", "Pilih Menu (1/2/3): ")
}

//fungsi untuk menambahkan pengguna/register
func menuRegister(arrayPengguna *arrID, nArrayPengguna *int) {
	var namaUser, Kategori, PassUser string
	var TargetUsser int

	fmt.Println("Nama Pengguna:")
	fmt.Scan(&namaUser)

	for i := 0; i < *nArrayPengguna; i++ {
		if namaUser == arrayPengguna[i].Username {
			fmt.Println("Nama sudah dipakai user lain, masukkan nama lain")
			fmt.Scan(&namaUser)
			i = -1 // reset ulang pengecekan
		}
	}

	fmt.Println("Password Akun:")
	fmt.Println("Pastikan password terdiri angka, simbol dan huruf besar")
	fmt.Scan(&PassUser)
	fmt.Println("Profesi (Pelajar/Mahasiswa/Pekerja, dll):")
	fmt.Scan(&Kategori)
	fmt.Println("Masukan Target Pencapaian Pendapatan (Input Target Hanya Angka): ")
	fmt.Scan(&TargetUsser)

	if *nArrayPengguna < NMAX {
		arrayPengguna[*nArrayPengguna].Username = namaUser
		arrayPengguna[*nArrayPengguna].PassUser = PassUser
		arrayPengguna[*nArrayPengguna].Kategori = Kategori
		arrayPengguna[*nArrayPengguna].Target = TargetUsser
		fmt.Println("Pengguna berhasil didaftarkan!")
		(*nArrayPengguna)++
	} else {
		fmt.Println("Kapasitas Pengguna Sudah Penuh")
	}
}

func menuLogin(arrayPengguna *arrID, nArrayPengguna *int) {
	fmt.Println("Masukkan Username:")
	var nama, password string
	fmt.Scan(&nama)
	fmt.Println("Masukan Password: ")
	fmt.Scan(&password)

	for i := 0; i < *nArrayPengguna; i++ {
		if nama == arrayPengguna[i].Username && password == arrayPengguna[i].PassUser {
			fmt.Println("Selamat! Anda berhasil login.")
			tampilkanDashboard(&arrayPengguna[i])

			return // login berhasil, keluar dari fungsi
		}
	}
	fmt.Println("Username atau Password salah.") // hanya tampil jika tidak ada yang cocok
}

// Fungsi untuk menambahkan data pendapatan
func tambahPendapatan(array *ID, n *int) {

	fmt.Println("\n--- Tambah Data Pendapatan ---")

	fmt.Println("Jenis Pendapatan: ")
	fmt.Println("1. Side Hustle")
	fmt.Println("2. Passive Income")
	fmt.Println("Pilih (1/2) :")

	var pilihan, jenisSideHustle, jenisPassive int
	if *n < NMAX {

		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Println("Jenis Side Hustle")
			fmt.Println("1. Freelance")
			fmt.Println("2. Konten Kreator")
			fmt.Println("3. Jualan Online")
			fmt.Println("Pilih (1/2/3): ")
			fmt.Scan(&jenisSideHustle)

			switch jenisSideHustle {
			case 1:
				array.dataPendapatan[*n].nama = "Freelance"
				KetSideHustle(array, n)
			case 2:
				array.dataPendapatan[*n].nama = "Konten Kreator"
				KetSideHustle(array, n)
			case 3:
				array.dataPendapatan[*n].nama = "Jualan Online"
				KetSideHustle(array, n)

			}

		case 2:
			fmt.Println("Jenis Passive Income")
			fmt.Println("1. Uang Sewa Lahan")
			fmt.Println("2. Royalti")
			fmt.Println("Pilih (1/2): ")
			fmt.Scan(&jenisPassive)

			switch jenisPassive {
			case 1:
				array.dataPendapatan[*n].nama = "Uang Sewa Lahan"
				KetSideHustle(array, n)
			case 2:
				array.dataPendapatan[*n].nama = "Royalti"
				KetSideHustle(array, n)
			}

		}
		(*n)++
	} else {
		fmt.Println("Data pendapatan sudah penuh.")
	}
}

// fungsi laporan keuangan lengkap dengan pilihan laporan bulanan dan tahunan
func laporan(array *ID, n *int) {
	var pilihanlap int
	var bulan, tahun int

	fmt.Println("===============")
	fmt.Println("1. Tampilkan Laporan Bulanan")
	fmt.Println("2. Tampilkan Laporan Tahunan")
	fmt.Print("Pilih (1/2): ")
	fmt.Scan(&pilihanlap)

	switch pilihanlap {
	case 1:
		fmt.Print("Masukkan Bulan (1-12): ")
		fmt.Scan(&bulan)
		fmt.Print("Masukkan Tahun: ")
		fmt.Scan(&tahun)
		lapBulanan(array, n, bulan, tahun)
	case 2:
		fmt.Print("Masukkan Tahun: ")
		fmt.Scan(&tahun)
		lapTahunan(array, n, tahun)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Binary search lower bound untuk tanggal
func binarySearchLowerBound(array *ID, n int, t, b, th int) int {
	low := 0
	high := n - 1
	result := -1
	for low <= high {
		mid := (low + high) / 2
		d := array.dataPendapatan[mid]

		if d.tahun > th ||
			(d.tahun == th && d.bulan > b) ||
			(d.tahun == th && d.bulan == b && d.tanggal >= t) {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}

// Binary search upper bound untuk tanggal
func binarySearchUpperBound(array *ID, n int, t, b, th int) int {
	low := 0
	high := n - 1
	result := -1
	for low <= high {
		mid := (low + high) / 2
		d := array.dataPendapatan[mid]

		if d.tahun < th ||
			(d.tahun == th && d.bulan < b) ||
			(d.tahun == th && d.bulan == b && d.tanggal <= t) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

// Sorting data pendapatan berdasarkan tanggal (tahun, bulan, tanggal)
func sortPendapatanByDate(array *ID, n *int) {
	for i := 0; i < *n-1; i++ {
		for j := i + 1; j < *n; j++ {
			d1 := array.dataPendapatan[i]
			d2 := array.dataPendapatan[j]

			if d1.tahun > d2.tahun ||
				(d1.tahun == d2.tahun && d1.bulan > d2.bulan) ||
				(d1.tahun == d2.tahun && d1.bulan == d2.bulan && d1.tanggal > d2.tanggal) {
				temp := array.dataPendapatan[i]
				array.dataPendapatan[i] = array.dataPendapatan[j]
				array.dataPendapatan[j] = temp
			}
		}
	}
}

func lapBulanan(array *ID, n *int, bulan, tahun int) {
	sortPendapatanByDate(array, n)
	fmt.Printf("\nLaporan Pendapatan Bulanan %02d-%d\n", bulan, tahun)

	start := binarySearchLowerBound(array, *n, 1, bulan, tahun)
	end := binarySearchUpperBound(array, *n, 31, bulan, tahun)

	if start == -1 || end == -1 || start > end {
		fmt.Println("Tidak ada data pendapatan di bulan ini.")
		return
	}

	var total float64 = 0
	for i := start; i <= end; i++ {
		d := array.dataPendapatan[i]
		fmt.Printf("- %s: %d (Tanggal %d)\n", d.nama, d.jumlah, d.tanggal)
		total += float64(d.jumlah)
	}
	fmt.Printf("Total Pendapatan Bulan %02d Tahun %d = %.2f\n", bulan, tahun, total)
}

func lapTahunan(array *ID, n *int, tahun int) {
	sortPendapatanByDate(array, n)
	fmt.Printf("\nLaporan Pendapatan Tahunan Tahun %d\n", tahun)

	start := binarySearchLowerBound(array, *n, 1, 1, tahun)
	end := binarySearchUpperBound(array, *n, 31, 12, tahun)

	if start == -1 || end == -1 || start > end {
		fmt.Println("Tidak ada data pendapatan di tahun ini.")
		return
	}

	var total float64 = 0
	for i := start; i <= end; i++ {
		d := array.dataPendapatan[i]
		fmt.Printf("- %s: %d (Tanggal %02d-%02d-%d)\n", d.nama, d.jumlah, d.tanggal, d.bulan, d.tahun)
		total += float64(d.jumlah)
	}
	fmt.Printf("Total Pendapatan Tahun %d = %.2f\n", tahun, total)
}

// Fungsi untuk menampilkan analisis (sementara kosong)
func TampilanAnalisis() {
	fmt.Println("Fitur analisis belum tersedia.")
}

// Fungsi untuk menampilkan progres pencapaian target
func TampilanProgres(array *ID, n *int) {
	fmt.Printf("Target Pendapatan Anda Bulan ini : %d\n", array.Target)
	fmt.Printf("Persentase Target Tercapai Bulan ini sebesar %.2f%%\n", hitungProgress(array, n))
}

func hitungProgress(array *ID, n *int) float64 {
	var total int
	for i := 0; i < *n; i++ {
		total += array.dataPendapatan[i].jumlah
	}
	if array.Target == 0 {
		return 0
	}
	return (float64(total) / float64(array.Target)) * 100
}

func KetSideHustle(array *ID, n *int) {
	fmt.Println("Masukan nominal:")
	fmt.Scan(&array.dataPendapatan[*n].jumlah)

	fmt.Println("Tanggal Pemasukan")
	fmt.Println("masukan Tanggal:")
	fmt.Scan(&array.dataPendapatan[*n].tanggal)

	fmt.Println("masukan bulan:")
	fmt.Scan(&array.dataPendapatan[*n].bulan)

	fmt.Println("masukan tahun:")
	fmt.Scan(&array.dataPendapatan[*n].tahun)

	fmt.Println("Deskripsi: ")
	fmt.Scan(&array.dataPendapatan[*n].deskripsi)

	fmt.Println("Pendapatan berhasil ditambahkan!")
}
