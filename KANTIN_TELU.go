package main

import (
	"fmt"
)

const NMAX int = 100

type pesan struct {
	pesan  [NMAX]kacow
	npesan int
}
type kacow struct {
	nama                       string
	harga                      int
	porsi                      int
	cek_login                  bool
	tot_pesan                  int
	tenant, username, password string
	cek_registrasi             bool
}

type deskripsi_pesanan struct {
	jumlah_pesanan, pilih_postent, pilih_tenant, porsi, tampilan_menu, input_budget int
	y_or_no                                                                         string
}

type tabint [NMAX]kacow

type tabPerubahan struct {
	nama  [NMAX]string //satu untuk tenant keberapa satu untuk menu kekeberapa-nya
	harga [NMAX]int
	sudah int
	tot   int
}

var j [NMAX]tabPerubahan
var pp pesan
var deskripsi deskripsi_pesanan

// program utama//W
func main() {
	var data tabint
	var str_n kacow
	var desk_pes deskripsi_pesanan
	Awal(data, desk_pes, &str_n)
}

// Awal adalah fungsi untuk menampilkan opsi awal
func Awal(A tabint, n deskripsi_pesanan, m *kacow) {
	var inputAwal int
	for inputAwal != 3 {
		fmt.Println("╔════════════════════════════════════════╗")
		fmt.Println("║                                        ║")
		fmt.Println("║             Selamat Datang             ║")
		fmt.Println("║                                        ║")
		fmt.Println("╠════════════════════════════════════════╣")
		fmt.Println("║          Pilih Opsi Kedatangan         ║")
		fmt.Println("║                                        ║")
		fmt.Println("║             1. ADMIN                   ║")
		fmt.Println("║             2. USER                    ║")
		fmt.Println("║             3. EXIT                    ║")
		fmt.Println("║                                        ║")
		fmt.Println("╚════════════════════════════════════════╝")
		fmt.Print("Ketikkan satu angka (1/2/3): ")
		fmt.Scan(&inputAwal)
		fmt.Println()
		if inputAwal == 1 {
			Login_admin(A, m)
		} else if inputAwal == 2 {
			login_mahasiswa(A, n, m)
		}
	}
}

func registrasi(m *kacow) {
	if m.cek_registrasi == false {
		fmt.Println("==================================================")
		fmt.Println("   Silahkan registrasi terlebih dahulu ^^")
		fmt.Println("==================================================")
		fmt.Print("🌟 Masukkan username anda: ")
		fmt.Scan(&m.username)

		fmt.Println()

		fmt.Print("🔑 Masukkan password anda: ")
		fmt.Scan(&m.password)
		fmt.Print("\033[H\033[2J")
	} else {

	}
}

// Login adalah fungsi untuk melakukan proses login//
func Login_admin(A tabint, n *kacow) {
	var login bool = false

	registrasi(n)
	fmt.Println("==================================================")
	fmt.Println("✅ Input username & password telah selesai!")
	fmt.Println("==================================================")

	var inputUserName, inputPassword string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&inputUserName)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&inputPassword)
	if n.username == inputUserName && n.password == inputPassword {
		login = true
		n.cek_registrasi = true
	} else {
		fmt.Println("**************************************************")
		fmt.Println("\033[31mPassword registrasi salah\033[0m") // Menambahkan warna merah pada teks
		fmt.Println("\033[33mSilahkan ulangi lagi!\033[0m")     // Menambahkan warna kuning pada teks
		fmt.Println("**************************************************")
	}
	if login {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Login berhasil")
		pilihan_admin(&A, n)
	} else {
		fmt.Println("Login tidak berhasil\n")
	}
}

func pilihan_admin(A *tabint, m *kacow) {
	var inputpilihanAdmin int

	option_admin()
	perubahan_menu(*A)
	fmt.Print("Ketikkan salah satu angka (1/2/3/4): ")
	fmt.Scan(&inputpilihanAdmin)

	if inputpilihanAdmin == 1 {
		//mengisi array baru
		fmt.Println("       📋 MENU ANDA TELAH DIINPUT  ")
		modifikasiMenu(*A)
	} else if inputpilihanAdmin == 2 {
		total_penghasilan(m)

	} else if inputpilihanAdmin == 3 {
		lihat_tenant_favorit(m, *A)
	}
}

func option_admin() {
	fmt.Println("╔═══════════════════════════════════════════╗")
	fmt.Println("║             🌟 PILIHAN ADMIN 🌟           ║")
	fmt.Println("╠═══════════════════════════════════════════╣")
	fmt.Println("║  1. ✏️  MODIFIKASI MENU                    ║")
	fmt.Println("║  2. 💰  HITUNG TOTAL PENGHASILAN          ║")
	fmt.Println("║  3. 📊  LIHAT TENANT FAVORIT              ║")
	fmt.Println("║  4. ❌  EXIT                              ║")
	fmt.Println("╚═══════════════════════════════════════════╝")
}

func modifikasiMenu(menu tabint) {
	perubahan_menu(menu)
	var pil1, pil2 int
	fmt.Print("\033[H\033[2J")
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║                                      ║")
	fmt.Println("║         📝SILAHKAN PILIH KANTIN      ║")
	fmt.Println("║                                      ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  1. Kantin Asrama Putra              ║")
	fmt.Println("║  2. Kantin Asrama Putri              ║")
	fmt.Println("║  3. Kantin Tmart                     ║")
	fmt.Println("║  4. Exit                             ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Print("Ketikkan salah satu angka (1/2/3/4): ")
	fmt.Scan(&pil2)
	if pil2 == 1 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("SILAHKAN PILIH TENANT")
		baca_menu_kacow()
		fmt.Print("ketikkan angka satu angka tenant (1/2/3/...) : ")
		fmt.Scan(&pil1)
		editMenu(pil1 - 1)
	} else if pil2 == 2 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("SILAHKAN PILIH TENANT")
		baca_menu_kacew()
		fmt.Print("ketikkan angka satu angka tenant (1/2/3/...) : ")
		fmt.Scan(&pil1)
		editMenu(pil1 + 4)
	} else if pil2 == 3 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("SILAHKAN PILIH TENANT")
		baca_menu_tmart()
		fmt.Print("ketikkan angka satu angka tenant (1/2/3/...) : ")
		fmt.Scan(&pil1)
		editMenu(pil1 + 10)
	} else {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Angka yang di masukkan salah, silah coba lagi ^^")
	}
}

func editMenu(pil int) {
	var pil2, harga int
	var pil3, nama string
	var yorno string
	var banyak_a int
	fmt.Print("\033[H\033[2J")
	fmt.Println("\t🌟 MENU TENANT 🌟")
	fmt.Println("NO   NAMA TENANT               HARGA")
	for i := 0; i < j[pil].tot; i++ {
		fmt.Printf("%-4d %-24s %-20d\n", i, j[pil].nama[i], j[pil].harga[i])
	}
	fmt.Println()
	fmt.Print("Ingin mengedit tampilan menu ? (Y/N) : ")
	fmt.Scan(&yorno)
	if yorno == "Y" || yorno == "y" {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════╗")
		fmt.Println("║                  🌟 PILIHAN MENU 🌟                  ║")
		fmt.Println("╠══════════════════════════════════════════════════════╣")
		fmt.Println("║  A. 🍔 Tambah menu                                   ║")
		fmt.Println("║  E. 📝 Edit menu                                     ║")
		fmt.Println("║  D. 🗑️ Delete                                         ║")
		fmt.Println("╚══════════════════════════════════════════════════════╝")
		fmt.Print("Silahkan pilih opsi A(tambah), E(edit), D(delete): ")
		fmt.Scan(&pil3)
		if pil3 == "e" || pil3 == "E" {
			var banyak_e int
			fmt.Print("Ingin merubah berapa banyak menu : ")
			fmt.Scan(&banyak_e)
			for i := 0; i < banyak_e; i++ {
				fmt.Print("Ketikkan angka menu yang ingin diubah :")
				fmt.Scan(&pil2)
				fmt.Print("Masukkan nama baru :")
				fmt.Scan(&nama)
				fmt.Println()
				fmt.Print("Masukkan harga baru: ")
				fmt.Scan(&harga)
				j[pil].nama[pil2] = nama
				j[pil].harga[pil2] = harga
			}
			fmt.Println("==========================================")
			for i := 0; i < 7+banyak_a; i++ {
				fmt.Println(j[pil].nama[i], j[pil].harga[i])
			}
		} else if pil3 == "a" || pil3 == "A" {
			fmt.Print("Ingin menambah berapa banyak menu : ")
			fmt.Scan(&banyak_a)
			for i := 0; i < banyak_a; i++ {
				fmt.Print("Masukkan nama :")
				fmt.Scan(&nama)
				fmt.Print("Masukkan harga : ")
				fmt.Scan(&harga)
				j[pil].nama[j[pil].tot] = nama
				j[pil].harga[j[pil].tot] = harga
				j[pil].tot++
				fmt.Println()
			}
			fmt.Println("==========================================")
			for i := 0; i < banyak_a+8; i++ {
				fmt.Println(j[pil].nama[i], j[pil].harga[i])
			}
			fmt.Println("==========================================")
		} else if pil3 == "d" || pil3 == "D" {
			var banyak_d int
			fmt.Print("Ingin menghapus berapa banyak menu : ")
			fmt.Scan(&banyak_d)
			fmt.Println()
			for Z := 0; Z < banyak_d; Z++ {
				fmt.Print("Ketikkan angka menu yang ingin dihapus :")
				fmt.Scan(&pil2)

				if pil2 < 0 || pil2 > j[pil].tot {
					fmt.Println("Menu tidak valid!")
				} else {
					for k := pil2; k < j[pil].tot; k++ {
						j[pil].nama[k] = j[pil].nama[k+1]
						j[pil].harga[k] = j[pil].harga[k+1]
					}
					j[pil].tot--

					fmt.Println("==========================================")
					for i := 0; i < j[pil].tot; i++ {
						fmt.Println(i, j[pil].nama[i], j[pil].harga[i])
					}
					fmt.Println("==========================================")

				}
			}
		}
		fmt.Println("menu anda saat ini telah diperbarui : ")
		fmt.Println()
	} else if yorno == "N" || yorno == "n" {

	} else {
		fmt.Println("Karakter yang di masukkan salah")
		editMenu(pil)
	}
}

func total_penghasilan(m *kacow) {
	var A tabint
	var n deskripsi_pesanan
	var total int
	var pendapatan_admin float64
	var pendapatan_tenant int
	if m.cek_login != true {
		fmt.Println("Anda harus melakukan pemesanan menu terlebih dahulu sebagai user")
		login_mahasiswa(A, n, m)
	} else {
		fmt.Println()
		fmt.Println("╔════════════════════════════════════════╗")
		fmt.Println("║          📜 Riwayat Transaksi          ║")
		fmt.Println("╠════════════════════════════════════════╣")
		j := 0
		pendapatan_admin = 0
		total = 0
		for j < pp.npesan {
			total = total + pp.pesan[j].porsi*pp.pesan[j].harga
			fmt.Println(j+1, "Pesanan saat ini ", pp.pesan[j].nama, " sebanyak ", pp.pesan[j].porsi)
			pendapatan_admin = pendapatan_admin + (float64(pp.pesan[j].porsi*pp.pesan[j].harga) * (0.25))
			pendapatan_tenant = total - int(pendapatan_admin)
			j++
		}
		fmt.Println("╠════════════════════════════════════════╣")
		fmt.Printf("║ Pendapatan Admin  : Rp%11.f ║\n", pendapatan_admin)
		fmt.Printf("║ Pendapatan Tenant : Rp%11d ║\n", pendapatan_tenant)
		fmt.Println("╚════════════════════════════════════════╝")
		fmt.Println()
		fmt.Println("🎉 Penghitungan total penghasilan selesai! 🎉")
	}
}

func lihat_tenant_favorit(m *kacow, A tabint) {
	var lanjut string
	fmt.Println()
	if !m.cek_login {
		fmt.Println("Silahkan Login Sebagai User Terlebih Dahulu")
		fmt.Print("Silahkan Masukkan Angka apapun untuk lanjut: ")
		fmt.Scan(&lanjut)
	} else {
		sorting_favorit()
		fmt.Println("Urutan tenant favorit :")
		fmt.Println()
		for j := 0; j < 14; j++ {
			fmt.Println(pp.pesan[j].tenant, "sudah dipesan sebanyak ", pp.pesan[j].tot_pesan)
		}
	}
}

func sorting_favorit() {
	var i, j, temp int
	var temp_tenant string
	for i = 1; i < 13; i++ {
		j = i
		temp = pp.pesan[i].tot_pesan
		temp_tenant = pp.pesan[i].tenant
		for j > 0 && pp.pesan[i-1].tot_pesan < temp {
			pp.pesan[j].tot_pesan = pp.pesan[j-1].tot_pesan
			pp.pesan[j].tenant = pp.pesan[j-1].tenant
			j--
		}
		pp.pesan[j].tot_pesan = temp
		pp.pesan[j].tenant = temp_tenant
	}
}

func login_mahasiswa(A tabint, n deskripsi_pesanan, m *kacow) {
	m.cek_login = true
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║                                      ║")
	fmt.Println("║     Silahkan pilih satu angka        ║")
	fmt.Println("║            (1/2/3/4)                 ║")
	fmt.Println("║                                      ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  1. Kantin Asrama Putra              ║")
	fmt.Println("║  2. Kantin Asrama Putri              ║")
	fmt.Println("║  3. Kantin Tmart                     ║")
	fmt.Println("║  4. EXIT                             ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Print("Ketikkan satu angka di sini: ")
	var input_angka_mahasiswa int

	//pilihan posisi kantin yang dituju//
	fmt.Scan(&input_angka_mahasiswa)
	fmt.Println()
	perubahan_menu(A)
	if input_angka_mahasiswa == 1 {
		kantin_asrama_putra(*m, A)
	} else if input_angka_mahasiswa == 2 {
		kantin_asrama_putri(*m, A)
	} else if input_angka_mahasiswa == 3 {
		kantin_tmart(*m, A)
	} else if input_angka_mahasiswa != 4 {
		fmt.Println()
		fmt.Println("Angka yang dimasukkan salah!\n")
		fmt.Println("Silahkan ulangi lagi\n")
		fmt.Println("==============================")
		login_mahasiswa(A, n, m)
	}
}

func kantin_asrama_putra(n kacow, menu tabint) {
	var deskripsi deskripsi_pesanan
	//jumlah_total_pesanan dari keseluruhan tenant//
	fmt.Print("Berapa jumlah pesanan makanan/minuman yang diinginkan: ")
	fmt.Scan(&deskripsi.jumlah_pesanan)
	fmt.Println()

	//pemesanan menu ini hanya bisa satu per satu menu dari setiap tenant//
	baca_menu_kacow()

	i := 0
	for i < deskripsi.jumlah_pesanan && deskripsi.pilih_postent < 6 {
		fmt.Print("Ketikkan satu angka disini: ")
		fmt.Scan(&deskripsi.pilih_postent)
		if deskripsi.pilih_postent == 1 {
			//proses simpan data menu//

			// proses cetak menu//
			fmt.Println("==================================")
			fmt.Println("         MENU DAPOER TRISNA       ")
			fmt.Println("==================================")
			fmt.Println()
		} else if deskripsi.pilih_postent == 2 {
			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("         	GUDEG YOGYA		       ")
			fmt.Println("==================================")
			fmt.Println()
		} else if deskripsi.pilih_postent == 3 {
			//proses simpan data menu//

			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("         	KEDAI JUICE		       ")
			fmt.Println("==================================")
			fmt.Println()
		} else if deskripsi.pilih_postent == 4 {
			//proses simpan data menu//

			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("         KANTIN ANUGRAH	       ")
			fmt.Println("==================================")
			fmt.Println()
		} else if deskripsi.pilih_postent == 5 {
			//proses simpan data menu//

			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("         		ASTRA		      ")
			fmt.Println("==================================")
			fmt.Println()
		}

		if deskripsi.pilih_postent < 6 {
			for k := 0; k < j[deskripsi.pilih_postent-1].tot; k++ {
				fmt.Println(k+1, j[deskripsi.pilih_postent-1].nama[k], j[deskripsi.pilih_postent-1].harga[k])
			}
			fmt.Println("==================================")

			//pengubahan tampilan menu dengan sorting dan serching//
			fmt.Println()
			fmt.Print("Ingin mengubah tampilan menu ?(N/Y) : ")
			fmt.Scan(&deskripsi.y_or_no)
			fmt.Println()

			if deskripsi.y_or_no == "Y" || deskripsi.y_or_no == "y" {
				tampilan_menu()
				fmt.Println()
				fmt.Print("Ketikkan satu angka disini (1/2/3..) : ")
				fmt.Scan(&deskripsi.tampilan_menu)
				fmt.Println()
				if deskripsi.tampilan_menu == 1 {
					var temp int
					temp = tampilan_menu_termurah_kacow(deskripsi.pilih_postent-1, menu)
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent-1].tot; i++ {
						if j[deskripsi.pilih_postent-1].harga[temp] == j[deskripsi.pilih_postent-1].harga[i] {
							fmt.Println(i+1, j[deskripsi.pilih_postent-1].nama[i], j[deskripsi.pilih_postent-1].harga[i])
						}
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 2 {
					var temp int
					temp = tampilan_menu_termahal_kacow(deskripsi.pilih_postent-1, menu)
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent-1].tot; i++ {
						if j[deskripsi.pilih_postent-1].harga[i] == j[deskripsi.pilih_postent-1].harga[temp] {
							fmt.Println(i+1, j[deskripsi.pilih_postent-1].nama[i], j[deskripsi.pilih_postent-1].harga[i])
						}
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 3 {
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent-1].tot; i++ {
						tampilan_menu_ascending_kacow(deskripsi.pilih_postent-1, &menu)
						fmt.Println(i+1, j[deskripsi.pilih_postent-1].nama[i], j[deskripsi.pilih_postent-1].harga[i])
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 4 {
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent-1].tot; i++ {
						tampilan_menu_descending_kacow(deskripsi.pilih_postent-1, &menu)
						fmt.Println(i+1, j[deskripsi.pilih_postent-1].nama[i], j[deskripsi.pilih_postent-1].harga[i])
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 5 {
					fmt.Println()
					fmt.Print("Isikan kisaran budget anda: ")
					fmt.Scan(&deskripsi.input_budget)
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent-1].tot; i++ {
						if j[deskripsi.pilih_postent-1].harga[i] <= deskripsi.input_budget {
							fmt.Println(i+1, j[deskripsi.pilih_postent-1].nama[i], j[deskripsi.pilih_postent-1].harga[i])
						}
					}
					fmt.Println("==============================")
				}
			} else {
				fmt.Println()
				fmt.Println("Silahkan lanjutkan pesanan anda ^^!")
			}

			//bagian format pemesanan//
			fmt.Print("Ingin memesan apa ? (1/2/3...):")
			fmt.Scan(&deskripsi.pilih_tenant)
			fmt.Print("berapa banyak porsi? (1/2/3...): ")
			fmt.Scan(&deskripsi.porsi)

			i = i + deskripsi.porsi
			pp.pesan[pp.npesan].nama = j[deskripsi.pilih_postent-1].nama[deskripsi.pilih_tenant-1]
			pp.pesan[pp.npesan].harga = j[deskripsi.pilih_postent-1].harga[deskripsi.pilih_tenant-1]
			pp.pesan[pp.npesan].porsi = deskripsi.porsi
			pp.npesan++

			//pesanan saat ini//
			fmt.Println()
			j := 0
			for j < pp.npesan {
				fmt.Println("Pesanan saat ini ", pp.pesan[j].nama, " sebanyak ", pp.pesan[j].porsi)
				j++
			}

			if i < deskripsi.jumlah_pesanan {
				fmt.Println("Silahkan pilih menu lagi ^^ !")
				baca_menu_kacow()
			} else if i == deskripsi.jumlah_pesanan {
				fmt.Println("Pesanan awal anda sebanyak ", deskripsi.jumlah_pesanan, " telah sesuai")
			}

			if deskripsi.pilih_postent == 1 {
				pp.pesan[0].tot_pesan = pp.pesan[0].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 2 {
				pp.pesan[1].tot_pesan = pp.pesan[1].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 3 {
				pp.pesan[2].tot_pesan = pp.pesan[2].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 4 {
				pp.pesan[3].tot_pesan = pp.pesan[3].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 5 {
				pp.pesan[4].tot_pesan = pp.pesan[4].tot_pesan + deskripsi.porsi
			}
		}
	}
}

func tampilan_menu() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                      PILIH TAMPILAN MENU                   ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║  1. 🌟 Cek harga termurah                                  ║")
	fmt.Println("║  2. 💰 Cek harga termahal                                  ║")
	fmt.Println("║  3. 📊 Urutkan dari harga termurah                         ║")
	fmt.Println("║  4. 📈 Urutkan dari harga termahal                         ║")
	fmt.Println("║  5. 🔍 Cari dengan harga budget                            ║")
	fmt.Println("║  6. ---EXIT---                                             ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}

func baca_menu_kacow() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         PILIH TENANT                       ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║  1. 🍜  DAPOER TRISHI - Masakan Tradisional                ║")
	fmt.Println("║  2. 🍛  GUDEG YOGYA - Gudeg dan Masakan Jawa               ║")
	fmt.Println("║  3. 🥤  KEDAI JUICE - Minuman Segar                        ║")
	fmt.Println("║  4. 🍪  KANTIN ANUGRAH - Makanan Ringan                    ║")
	fmt.Println("║  5. 🍔  ASTRA - Makanan Cepat Saji                         ║")
	fmt.Println("║  6. ---EXIT---                                             ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}

func kantin_asrama_putri(n kacow, A tabint) {
	var deskripsi deskripsi_pesanan
	//jumlah_total_pesanan dari keseluruhan tenant//
	fmt.Print("Berapa jumlah pesanan makanan/minuman yang diinginkan: ")
	fmt.Scan(&deskripsi.jumlah_pesanan)
	fmt.Println()

	baca_menu_kacew()

	i := 0
	for i < deskripsi.jumlah_pesanan && deskripsi.pilih_postent < 7 {
		fmt.Print("Ketikkan satu angka disini: ")
		fmt.Scan(&deskripsi.pilih_postent)
		fmt.Println()
		if deskripsi.pilih_postent == 1 {
			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("         MENU DAPUR D2       ")
			fmt.Println("==================================")
			fmt.Println()

		} else if deskripsi.pilih_postent == 2 {
			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("         MENU KEDAI SANTI	      ")
			fmt.Println("==================================")
			fmt.Println()

		} else if deskripsi.pilih_postent == 3 {
			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("         MENU KEDAI RUMAHAN      ")
			fmt.Println("==================================")
			fmt.Println()

		} else if deskripsi.pilih_postent == 4 {
			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("         MENU KANTIN ALWI	      ")
			fmt.Println("==================================")
			fmt.Println()

		} else if deskripsi.pilih_postent == 5 {
			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("         MENU KANTIN KARI	      ")
			fmt.Println("==================================")
			fmt.Println()

		} else if deskripsi.pilih_postent == 6 {
			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("        	 MENU KRIUW		      ")
			fmt.Println("==================================")
			fmt.Println()

		}
		if deskripsi.pilih_postent < 7 {
			deskripsi.pilih_postent = deskripsi.pilih_postent + 4
			for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
				fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
			}
			fmt.Println("==================================")

			//pengubahan tampilan menu dengan sorting dan serching//
			fmt.Println()
			fmt.Print("Ingin mengubah tampilan menu ?(N/Y) : ")
			fmt.Scan(&deskripsi.y_or_no)
			fmt.Println()

			if deskripsi.y_or_no == "Y" || deskripsi.y_or_no == "y" {
				tampilan_menu()
				fmt.Println()
				fmt.Print("Ketikkan satu angka disini : (1/2/3..) : ")
				fmt.Scan(&deskripsi.tampilan_menu)
				fmt.Println()
				if deskripsi.tampilan_menu == 1 {
					var temp int
					temp = tampilan_menu_termurah_kacow(deskripsi.pilih_postent, A)
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
						if j[deskripsi.pilih_postent].harga[temp] == j[deskripsi.pilih_postent].harga[i] {
							fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
						}
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 2 {
					var temp int
					temp = tampilan_menu_termahal_kacow(deskripsi.pilih_postent, A)
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
						if j[deskripsi.pilih_postent].harga[i] == j[deskripsi.pilih_postent].harga[temp] {
							fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
						}
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 3 {
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
						tampilan_menu_ascending_kacow(deskripsi.pilih_postent, &A)
						fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 4 {
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
						tampilan_menu_descending_kacow(deskripsi.pilih_postent, &A)
						fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 5 {
					fmt.Println()
					fmt.Print("Isikan kisaran budget anda: ")
					fmt.Scan(&deskripsi.input_budget)
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
						if j[deskripsi.pilih_postent].harga[i] <= deskripsi.input_budget {
							fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
						}
					}
					fmt.Println("==============================")
				}
			} else {
				fmt.Println()
				fmt.Println("Silahkan lanjutkan pesanan anda ^^!")
			}

			//bagian format pemesanan//
			fmt.Print("Ingin memesan apa ? (1/2/3...):")
			fmt.Scan(&deskripsi.pilih_tenant)
			fmt.Print("berapa banyak porsi? (1/2/3...): ")
			fmt.Scan(&deskripsi.porsi)

			if deskripsi.pilih_postent == 1 {
				pp.pesan[5].tot_pesan = pp.pesan[5].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 2 {
				pp.pesan[6].tot_pesan = pp.pesan[6].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 3 {
				pp.pesan[7].tot_pesan = pp.pesan[7].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 4 {
				pp.pesan[8].tot_pesan = pp.pesan[8].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 5 {
				pp.pesan[9].tot_pesan = pp.pesan[9].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 6 {
				pp.pesan[10].tot_pesan = pp.pesan[10].tot_pesan + deskripsi.porsi
			}

			i = i + deskripsi.porsi
			pp.pesan[pp.npesan].nama = j[deskripsi.pilih_postent].nama[deskripsi.pilih_tenant-1]
			pp.pesan[pp.npesan].harga = j[deskripsi.pilih_postent].harga[deskripsi.pilih_tenant-1]
			pp.pesan[pp.npesan].porsi = deskripsi.porsi
			pp.npesan++

			//pesanan saat ini//
			fmt.Println()
			j := 0
			for j < pp.npesan {
				fmt.Println("Pesanan saat ini ", pp.pesan[j].nama, " sebanyak ", pp.pesan[j].porsi)
				j++
			}

			if i < deskripsi.jumlah_pesanan {
				fmt.Println("Silahkan pilih menu lagi ^^ !")
				baca_menu_kacew()
			} else if i == deskripsi.jumlah_pesanan {
				fmt.Println("Pesanan awal anda sebanyak ", deskripsi.jumlah_pesanan, " telah sesuai")
			}
		}
	}
}

func baca_menu_kacew() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                          PILIH TENANT                      ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║  1. 🍜  DAPUR D2 - Citarasa Autentik!                      ║")
	fmt.Println("║  2. ☕  KEDAI SANTI - Aroma Khas Nusantara!                ║")
	fmt.Println("║  3. 🏠  KEDAI RUMAHAN - Suasana Hangat di Tengah Kota!     ║")
	fmt.Println("║  4. 🍱  KANTIN ALWI - Pilihan Gizi Terbaik untukmu!        ║")
	fmt.Println("║  5. 🌶️  KANTIN KARI - Pedasnya Bikin Ketagihan!           ║")
	fmt.Println("║  6. 🍔  KRIUWW - Kelezatan Tanpa Batas!                    ║")
	fmt.Println("║  7. ---EXIT---                                             ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}

func kantin_tmart(n kacow, A tabint) {
	var deskripsi deskripsi_pesanan
	//jumlah_total_pesanan dari keseluruhan tenant//
	fmt.Print("Berapa jumlah pesanan makanan/minuman yang diinginkan: ")
	fmt.Scan(&deskripsi.jumlah_pesanan)
	fmt.Println()

	baca_menu_tmart()

	i := 0
	for i < deskripsi.jumlah_pesanan && deskripsi.pilih_postent < 4 {
		fmt.Print("Ketikkan satu angka disini: ")
		fmt.Scan(&deskripsi.pilih_postent)
		fmt.Println()
		if deskripsi.pilih_postent == 1 {
			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("       MENU DAPUR BAHARI	     ")
			fmt.Println("==================================")
			fmt.Println()

		} else if deskripsi.pilih_postent == 2 {
			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("       MENU KEDAI WIWIT	     ")
			fmt.Println("==================================")
			fmt.Println()

		} else if deskripsi.pilih_postent == 3 {
			// proses cetak menu//

			fmt.Println("==================================")
			fmt.Println("       MENU KEDAI APIW	     ")
			fmt.Println("==================================")
			fmt.Println()

		}
		if deskripsi.pilih_postent < 4 {
			deskripsi.pilih_postent = deskripsi.pilih_postent + 10
			for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
				fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
			}
			fmt.Println("==================================")

			//pengubahan tampilan menu dengan sorting dan serching//
			fmt.Println()
			fmt.Print("Ingin mengubah tampilan menu ?(N/Y) : ")
			fmt.Scan(&deskripsi.y_or_no)
			fmt.Println()

			if deskripsi.y_or_no == "Y" || deskripsi.y_or_no == "y" {
				tampilan_menu()
				fmt.Println()
				fmt.Print("Ketikkan satu angka disini : (1/2/3..) : ")
				fmt.Scan(&deskripsi.tampilan_menu)
				fmt.Println()
				if deskripsi.tampilan_menu == 1 {
					var temp int
					temp = tampilan_menu_termurah_kacow(deskripsi.pilih_postent, A)
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
						if j[deskripsi.pilih_postent].harga[temp] == j[deskripsi.pilih_postent].harga[i] {
							fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
						}
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 2 {
					var temp int
					temp = tampilan_menu_termahal_kacow(deskripsi.pilih_postent, A)
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
						if j[deskripsi.pilih_postent].harga[i] == j[deskripsi.pilih_postent].harga[temp] {
							fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
						}
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 3 {
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
						tampilan_menu_ascending_kacow(deskripsi.pilih_postent, &A)
						fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 4 {
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
						tampilan_menu_descending_kacow(deskripsi.pilih_postent, &A)
						fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
					}
					fmt.Println("==============================")
				} else if deskripsi.tampilan_menu == 5 {
					fmt.Println()
					fmt.Print("Isikan kisaran budget anda: ")
					fmt.Scan(&deskripsi.input_budget)
					fmt.Println("==============================")
					for i := 0; i < j[deskripsi.pilih_postent].tot; i++ {
						if j[deskripsi.pilih_postent].harga[i] <= deskripsi.input_budget {
							fmt.Println(i+1, j[deskripsi.pilih_postent].nama[i], j[deskripsi.pilih_postent].harga[i])
						}
					}
					fmt.Println("==============================")
				}
			} else {
				fmt.Println()
				fmt.Println("Silahkan lanjutkan pesanan anda ^^!")
			}

			//bagian format pemesanan//
			fmt.Print("Ingin memesan apa ? (1/2/3...):")
			fmt.Scan(&deskripsi.pilih_tenant)
			fmt.Print("berapa banyak porsi? (1/2/3...): ")
			fmt.Scan(&deskripsi.porsi)

			if deskripsi.pilih_postent == 1 {
				pp.pesan[11].tot_pesan = pp.pesan[11].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 2 {
				pp.pesan[12].tot_pesan = pp.pesan[12].tot_pesan + deskripsi.porsi
			} else if deskripsi.pilih_postent == 3 {
				pp.pesan[13].tot_pesan = pp.pesan[13].tot_pesan + deskripsi.porsi
			}

			i = i + deskripsi.porsi
			pp.pesan[pp.npesan].nama = j[deskripsi.pilih_postent].nama[deskripsi.pilih_tenant-1]
			pp.pesan[pp.npesan].harga = j[deskripsi.pilih_postent].harga[deskripsi.pilih_tenant-1]
			pp.pesan[pp.npesan].porsi = deskripsi.porsi
			pp.npesan++

			//pesanan saat ini//
			fmt.Println()
			j := 0
			for j < pp.npesan {
				fmt.Println("Pesanan saat ini ", pp.pesan[j].nama, " sebanyak ", pp.pesan[j].porsi)
				j++
			}
			if i < deskripsi.jumlah_pesanan {
				fmt.Println("Silahkan pilih menu lagi ^^ !")
				baca_menu_tmart()
			} else if i == deskripsi.jumlah_pesanan {
				fmt.Println("Pesanan awal anda sebanyak ", deskripsi.jumlah_pesanan, " telah sesuai")
			}
		}
	}
}

func perubahan_menu(A tabint) {
	pp.pesan[0].tenant = "DAPOER TRISNA"
	pp.pesan[1].tenant = "GUDEG YOGYA"
	pp.pesan[2].tenant = "KEDAI JUICE"
	pp.pesan[3].tenant = "KANTIN ANUGRAH"
	pp.pesan[4].tenant = "ASTRA"
	pp.pesan[5].tenant = "DAPUR D2"
	pp.pesan[6].tenant = "KEDAI SANTI"
	pp.pesan[7].tenant = "KEDAI RUMAHAN"
	pp.pesan[8].tenant = "KANTIN ALWI"
	pp.pesan[9].tenant = "KANTIN KARI"
	pp.pesan[10].tenant = "KANTIN KRIUW"
	pp.pesan[11].tenant = "DAPUR BAHARI"
	pp.pesan[12].tenant = "KEDAI WIWIT"
	pp.pesan[13].tenant = "KEDAI APIW"

	if j[0].sudah == 0 { //kantin cowok
		A[0].nama = "Nasi_goreng_ayam_serundeng" //trisna
		A[0].harga = 20000
		j[0].nama[0] = A[0].nama
		j[0].harga[0] = A[0].harga

		A[1].nama = "Nasi_goreng_ayam_komplit"
		A[1].harga = 20000
		j[0].nama[1] = A[1].nama
		j[0].harga[1] = A[1].harga

		A[2].nama = "Nasi_gila"
		A[2].harga = 20000
		j[0].nama[2] = A[2].nama
		j[0].harga[2] = A[2].harga

		A[3].nama = "Chicken_cordon_blue"
		A[3].harga = 23000
		j[0].nama[3] = A[3].nama
		j[0].harga[3] = A[3].harga

		A[4].nama = "Nasi_bakar"
		A[4].harga = 20000
		j[0].nama[4] = A[4].nama
		j[0].harga[4] = A[4].harga

		A[5].nama = "Nasi_goreng_seafood"
		A[5].harga = 23000
		j[0].nama[5] = A[5].nama
		j[0].harga[5] = A[5].harga

		A[6].nama = "Rice_bowl"
		A[6].harga = 20000
		j[0].nama[6] = A[6].nama
		j[0].harga[6] = A[6].harga

		A[7].nama = "Nasi_goreng_omllete"
		A[7].harga = 20000
		j[0].nama[7] = A[7].nama
		j[0].harga[7] = A[7].harga
		j[0].tot = j[0].tot + 8
		j[0].sudah++
	}

	if j[1].sudah == 0 {
		A[0].nama = "Nasi_gudeg" //gudeg
		A[0].harga = 20000
		j[1].nama[0] = A[0].nama
		j[1].harga[0] = A[0].harga
		j[1].tot++

		A[1].nama = "Ayam_penyet"
		A[1].harga = 15000
		j[1].nama[1] = A[1].nama
		j[1].harga[1] = A[1].harga
		j[1].tot++

		A[2].nama = "Nasi_goreng_hongkong"
		A[2].harga = 20000
		j[1].nama[2] = A[2].nama
		j[1].harga[2] = A[2].harga
		j[1].tot++

		A[3].nama = "Nasi_timbel"
		A[3].harga = 18000
		j[1].nama[3] = A[3].nama
		j[1].harga[3] = A[3].harga
		j[1].tot++

		A[4].nama = "Soto_ayam_dan_nasi"
		A[4].harga = 20000
		j[1].nama[4] = A[4].nama
		j[1].harga[4] = A[4].harga
		j[1].tot++

		A[5].nama = "Ayam_teriyaki"
		A[5].harga = 18000
		j[1].nama[5] = A[5].nama
		j[1].harga[5] = A[5].harga
		j[1].tot++

		A[6].nama = "Rice_bowl"
		A[6].harga = 20000
		j[1].nama[6] = A[6].nama
		j[1].harga[6] = A[6].harga
		j[1].tot++
		j[1].sudah++
	}

	if j[2].sudah == 0 {
		A[0].nama = "Jus_pepaya" //kedai juice
		A[0].harga = 8000
		j[2].nama[0] = A[0].nama
		j[2].harga[0] = A[0].harga
		j[2].tot++

		A[1].nama = "Jus_tomat"
		A[1].harga = 8000
		j[2].nama[1] = A[1].nama
		j[2].harga[1] = A[1].harga
		j[2].tot++

		A[2].nama = "Jus_mangga"
		A[2].harga = 10000
		j[2].nama[2] = A[2].nama
		j[2].harga[2] = A[2].harga
		j[2].tot++

		A[3].nama = "Jus_alpukat"
		A[3].harga = 13000
		j[2].nama[3] = A[3].nama
		j[2].harga[3] = A[3].harga
		j[2].tot++

		A[4].nama = "Es_teh"
		A[4].harga = 5000
		j[2].nama[4] = A[4].nama
		j[2].harga[4] = A[4].harga
		j[2].tot++

		A[5].nama = "Kopi_ABC"
		A[5].harga = 5000
		j[2].nama[5] = A[5].nama
		j[2].harga[5] = A[5].harga
		j[2].tot++

		A[6].nama = "Pop_ice"
		A[6].harga = 10000
		j[2].nama[6] = A[6].nama
		j[2].harga[6] = A[6].harga
		j[2].tot++
		j[2].sudah++
	}

	if j[3].sudah == 0 {
		A[0].nama = "Nasi_burger_isi_tongkol_rica"
		A[0].harga = 9000
		j[3].nama[0] = A[0].nama
		j[3].harga[0] = A[0].harga
		j[3].tot++

		A[1].nama = "Nasi_cumi/udang/kerang_asam_manis"
		A[1].harga = 29000
		j[3].nama[1] = A[1].nama
		j[3].harga[1] = A[1].harga
		j[3].tot++

		A[2].nama = "Nasi_dan_ayam_kriwil"
		A[2].harga = 20000
		j[3].nama[2] = A[2].nama
		j[3].harga[2] = A[2].harga
		j[3].tot++

		A[3].nama = "Nasi_goreng_dan_ayam_goreng/gulung/kriwil"
		A[3].harga = 20000
		j[3].nama[3] = A[3].nama
		j[3].harga[3] = A[3].harga
		j[3].tot++

		A[4].nama = "Indomie_rebus_aneka_rasa"
		A[4].harga = 15000
		j[3].nama[4] = A[4].nama
		j[3].harga[4] = A[4].harga
		j[3].tot++

		A[5].nama = "Nasi_dan_gule_daging_sapi"
		A[5].harga = 27000
		j[3].nama[5] = A[5].nama
		j[3].harga[5] = A[5].harga
		j[3].tot++

		A[6].nama = "Nasi_dan_soto/telur"
		A[6].harga = 19000
		j[3].nama[6] = A[6].nama
		j[3].harga[6] = A[6].harga
		j[3].tot++
		j[3].sudah++
	}

	if j[4].sudah == 0 {
		A[0].nama = "Mie_nyemek"
		A[0].harga = 17000
		j[4].nama[0] = A[0].nama
		j[4].harga[0] = A[0].harga
		j[4].tot++

		A[1].nama = "Mie_goreng"
		A[1].harga = 18000
		j[4].nama[1] = A[1].nama
		j[4].harga[1] = A[1].harga
		j[4].tot++

		A[2].nama = "Nasi_dan_ayam_crispy_geprek/katsu"
		A[2].harga = 19000
		j[4].nama[2] = A[2].nama
		j[4].harga[2] = A[2].harga
		j[4].tot++

		A[3].nama = "Ayam_goreng_dan_nasi"
		A[3].harga = 20000
		j[4].nama[3] = A[3].nama
		j[4].harga[3] = A[3].harga
		j[4].tot++

		A[4].nama = "Tongseng_ayam_dan_nasi"
		A[4].harga = 20000
		j[4].nama[4] = A[4].nama
		j[4].harga[4] = A[4].harga
		j[4].tot++

		A[5].nama = "Nasi_goreng_telur"
		A[5].harga = 16000
		j[4].nama[5] = A[5].nama
		j[4].harga[5] = A[5].harga
		j[4].tot++

		A[6].nama = "Nasi_goreng_katsu"
		A[6].harga = 22000
		j[4].nama[6] = A[6].nama
		j[4].harga[6] = A[6].harga
		j[4].tot++
		j[4].sudah++
	}

	if j[5].sudah == 0 { //kantin cewek
		A[0].nama = "Tongseng_Ayam"
		A[0].harga = 18000

		A[1].nama = "Korean_Chicken"
		A[1].harga = 20000

		A[2].nama = "Karage_Cabe_Garam"
		A[2].harga = 20000

		A[3].nama = "Karage_Tereyaki"
		A[3].harga = 20000

		A[4].nama = "Karage_Asam_Manis"
		A[4].harga = 20000

		A[5].nama = "Karage_Mozarella"
		A[5].harga = 22000

		A[6].nama = "Karage_Black_Pepper"
		A[6].harga = 20000

		for i := 0; i < 7; i++ {
			j[5].nama[i] = A[i].nama
			j[5].harga[i] = A[i].harga
			j[5].tot++
		}
		j[5].sudah++
	}
	if j[6].sudah == 0 {
		A[0].nama = "Katsu_sambal_matah"
		A[0].harga = 19000

		A[1].nama = "Katsu_tereyaki"
		A[1].harga = 18000

		A[2].nama = "Nasi_goreng_telur_sosis"
		A[2].harga = 15000

		A[3].nama = "Nasi_goreng_katsu"
		A[3].harga = 20000

		A[4].nama = "Spageti_katsu"
		A[4].harga = 18000

		A[5].nama = "Spageti_oglio_lio"
		A[5].harga = 15000

		A[6].nama = "Cuanki_bandung"
		A[6].harga = 12000
		for i := 0; i < 7; i++ {
			j[6].nama[i] = A[i].nama
			j[6].harga[i] = A[i].harga
			j[6].tot++
		}
		j[6].sudah++
	}

	if j[7].sudah == 0 {
		A[0].nama = "Rice_bowl"
		A[0].harga = 15000

		A[1].nama = "Ayam_madu"
		A[1].harga = 18000

		A[2].nama = "Mie_tektek"
		A[2].harga = 15000

		A[3].nama = "Mie_nyemek"
		A[3].harga = 20000

		A[4].nama = "Ayam_geprek"
		A[4].harga = 17000

		A[5].nama = "Ayam_bakar"
		A[5].harga = 18000

		A[6].nama = "Ikan_goreng"
		A[6].harga = 18000
		for i := 0; i < 7; i++ {
			j[7].nama[i] = A[i].nama
			j[7].harga[i] = A[i].harga
			j[7].tot++
		}
		j[7].sudah++
	}
	if j[8].sudah == 0 {
		A[0].nama = "Nasi_panggang_katsu_moza"
		A[0].harga = 23000

		A[1].nama = "Nasi_katsu_sambal_matah"
		A[1].harga = 22000

		A[2].nama = "Nasi_daun_jeruk_ayam_geprek"
		A[2].harga = 22000

		A[3].nama = "Spageti_bolognese_katsu"
		A[3].harga = 21000

		A[4].nama = "Nasi_goreng_katsu_moza"
		A[4].harga = 26000

		A[5].nama = "Ayam_sambal_kemangi"
		A[5].harga = 22000

		A[6].nama = "Cordon_blue"
		A[6].harga = 22000
		for i := 0; i < 7; i++ {
			j[8].nama[i] = A[i].nama
			j[8].harga[i] = A[i].harga
			j[8].tot++
		}
		j[8].sudah++
	}
	if j[9].sudah == 0 {
		A[0].nama = "Nasi_kari_sapi"
		A[0].harga = 23000

		A[1].nama = "Sup_rempah _api"
		A[1].harga = 23000

		A[2].nama = "Soto_sapi"
		A[2].harga = 23000

		A[3].nama = "Soto_kikil"
		A[3].harga = 23000

		A[4].nama = "Soto_ayam"
		A[4].harga = 20000

		A[5].nama = "Nasi_capcay"
		A[5].harga = 21000

		A[6].nama = "Cumi_cabe_ijo"
		A[6].harga = 23000
		for i := 0; i < 7; i++ {
			j[9].nama[i] = A[i].nama
			j[9].harga[i] = A[i].harga
			j[9].tot++
		}
		j[9].sudah++
	}
	if j[10].sudah == 0 {
		A[0].nama = "Gyudon_buri"
		A[0].harga = 25000

		A[1].nama = "Gyudon_sambal"
		A[1].harga = 25000

		A[2].nama = "Dori_matah"
		A[2].harga = 25000

		A[3].nama = "Dori_cabe_garam"
		A[3].harga = 25000

		A[4].nama = "Tempura_sambal_matah"
		A[4].harga = 25000

		A[5].nama = "Tempura_sadass"
		A[5].harga = 25000

		A[6].nama = "Nasgor_sambal_ijo"
		A[6].harga = 20000
		for i := 0; i < 7; i++ {
			j[10].nama[i] = A[i].nama
			j[10].harga[i] = A[i].harga
			j[10].tot++
		}
		j[10].sudah++
	}
	if j[11].sudah == 0 { //kantin tmart
		A[0].nama = "Batagor_kering"
		A[0].harga = 15000

		A[1].nama = "Batagor_kuah"
		A[1].harga = 15000

		A[2].nama = "Pentol_kuah"
		A[2].harga = 12000

		A[3].nama = "Mie_ayam"
		A[3].harga = 16000

		A[4].nama = "Indomie_bakso"
		A[4].harga = 16000

		A[5].nama = "Nasgor_katsu"
		A[5].harga = 16000

		A[6].nama = "Cilok_kuah"
		A[6].harga = 11000
		for i := 0; i < 7; i++ {
			j[11].nama[i] = A[i].nama
			j[11].harga[i] = A[i].harga
			j[11].tot++
		}
		j[11].sudah++
	}
	if j[12].sudah == 0 {
		A[0].nama = "Ayam_kari"
		A[0].harga = 21000

		A[1].nama = "Fish_roll"
		A[1].harga = 23000

		A[2].nama = "Suki_kuah"
		A[2].harga = 22000

		A[3].nama = "Soto_ayam"
		A[3].harga = 17000

		A[4].nama = "Ayam_bakar"
		A[4].harga = 18000

		A[5].nama = "Nasi_cumi"
		A[5].harga = 21000

		A[6].nama = "Ayam_daun_jeruk"
		A[6].harga = 18000
		for i := 0; i < 7; i++ {
			j[12].nama[i] = A[i].nama
			j[12].harga[i] = A[i].harga
			j[12].tot++
		}
		j[12].sudah++
	}
	if j[13].sudah == 0 {
		A[0].nama = "Soto_ayam"
		A[0].harga = 18000

		A[1].nama = "Sup_kare"
		A[1].harga = 20000

		A[2].nama = "Bakso_malang"
		A[2].harga = 20000

		A[3].nama = "Cuanki_kuah"
		A[3].harga = 20000

		A[4].nama = "Mie_pangsit"
		A[4].harga = 20000

		A[5].nama = "Mie_yamin_manis"
		A[5].harga = 22000

		A[6].nama = "Mie_yamin_asin"
		A[6].harga = 20000
		for i := 0; i < 7; i++ {
			j[13].nama[i] = A[i].nama
			j[13].harga[i] = A[i].harga
			j[13].tot++
		}
		j[13].sudah++
	}
}

func baca_menu_tmart() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                        🌟 PILIH TENANT                     ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║  1. 🍳 DAPUR BAHARI - Sajian Makanan Lezat                 ║")
	fmt.Println("║  2. ☕ KEDAI WIWIT - Aroma Menu Terbaik                    ║")
	fmt.Println("║  3. 🍜 KEDAI APIW - Pilihan Kuliner Nusantara              ║")
	fmt.Println("║  7. ---EXIT---                                             ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}

func tampilan_menu_termurah_kacow(pil int, A tabint) int {
	var idx_min int
	idx_min = 0
	for i := 0; i < 7; i++ {
		if j[pil].harga[idx_min] > j[pil].harga[i] {
			idx_min = i
		}
	}
	return idx_min
}

func tampilan_menu_termahal_kacow(pil int, A tabint) int {
	var idx_max int
	idx_max = 0
	for i := 0; i < 7; i++ {
		if j[pil].harga[idx_max] < j[pil].harga[i] {
			idx_max = i
		}
	}
	return idx_max
}

func tampilan_menu_ascending_kacow(pil int, A *tabint) {
	var pass, ji, idx, temp int
	var temp2 string
	for pass = 1; pass < j[pil].tot; pass++ {
		idx = pass - 1
		for ji = pass; ji < j[pil].tot; ji++ {
			if j[pil].harga[idx] > j[pil].harga[ji] {
				idx = ji
			}
		}
		temp = j[pil].harga[idx]
		j[pil].harga[idx] = j[pil].harga[pass-1]
		j[pil].harga[pass-1] = temp

		temp2 = j[pil].nama[idx]
		j[pil].nama[idx] = j[pil].nama[pass-1]
		j[pil].nama[pass-1] = temp2
	}
}

func tampilan_menu_descending_kacow(pil int, A *tabint) {
	var pass, ji, idx, temp int
	var temp2 string
	for pass = 1; pass < j[pil].tot; pass++ {
		idx = pass - 1
		for ji = pass; ji < j[pil].tot; ji++ {
			if j[pil].harga[idx] < j[pil].harga[ji] {
				idx = ji
			}
		}
		temp = j[pil].harga[idx]
		j[pil].harga[idx] = j[pil].harga[pass-1]
		j[pil].harga[pass-1] = temp

		temp2 = j[pil].nama[idx]
		j[pil].nama[idx] = j[pil].nama[pass-1]
		j[pil].nama[pass-1] = temp2
	}
}
