// smile

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	//	"time"
)

const version = "0.5.0"

var verbose bool

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func checkBool(text string) bool {
	var condition bool

	if text == "sim" || text == "yes" {
		condition = true
	}

	return condition
}
func checkResponse(response bool) string {

	if response == true {
		return "yes"
	} else {
		return "no"
	}
}

func execute(cmd string) []byte {

	cmdLine := strings.Fields(cmd)
	command := cmdLine[0]
	parameters := cmdLine[1:len(cmdLine)]

	out, err := exec.Command(command, parameters...).Output()

	check(err)
	if verbose == true {
		fmt.Printf("%s", out)
	}
	return out
}

type ConnectionProfile struct {
	wifiInterface    string
	connectionType   string
	wifiSecurityType string
	essid            string
	ipMode           string
	wifiPassword     string
	hidden           bool
}

func executeInArchChroot(cmd string) []byte {



        out, err := exec.Command("bash", "-c","arch-chroot '/mnt' '/bin/bash' -c '" + cmd + "'").Output()

	check(err)

	return out
				
}

func (connProfile *ConnectionProfile) writeWifiConfigToFile(destinationFolder string) {

	const description = "Description=Configuration file created by smile " + version
	const descWifiInterface = "Interface="
	const descConnectionType = "Connection="
	const descWifiSecurityType = "Security="
	const descEssid = "ESSID="
	const descIpMode = "IP="
	const descWifiPass = "Key="
	const descHidden = "Hidden="

	data := []byte(description + "\n" +

		descWifiInterface + connProfile.wifiInterface + "\n" +

		descConnectionType + connProfile.connectionType + "\n" +

		descWifiSecurityType + connProfile.wifiSecurityType + "\n" +

		descEssid + connProfile.essid + "\n" +

		descIpMode + connProfile.ipMode + "\n" +

		descWifiPass + connProfile.wifiPassword + "\n" +

		descHidden + checkResponse(connProfile.hidden))

	err := ioutil.WriteFile(destinationFolder+"/"+connProfile.essid, data, 0777)
	check(err)

}
func detectNetwork() []string {

	cmdOut := execute("ip link")

	r, _ := regexp.Compile("[a-z]{3}[0-9][a-z][0-9]")

	regex := r.FindAllString(string(cmdOut), -1)
	return regex
}
func detectDevice() []string {

	cmdOut := execute("lsblk")
	r, _ := regexp.Compile("[s]{1}[d]{1}[a-z]{1}[ ]{1}")
	regex := r.FindAllString(string(cmdOut), -1)
	return regex
}

//func rankMirrors(){

//	fmt.Println("Será efetuado o processo de ranking dos servidores mais usados. Isso pode demorar alguns minutos")
//	_ = execute ("cp /etc/pacman.d/mirrorlist /etc/pacman.d/mirrorlist.backup")
//	_ = execute ("rankmirrors -n 5 /etc/pacman.d/mirrorlist.backup > /etc/pacman.d/mirrorlist")
//}

func (connProfile *ConnectionProfile) printConnectionProfile() {
	fmt.Printf("%s", execute("clear"))
	fmt.Println("\t\t\t\t\tOs dados informados foram: \n\n" +

		"\t\t\t\t\tInterface de rede: " + connProfile.wifiInterface + "\n\n" +

		"\t\t\t\t\tTipo de Conexão: " + connProfile.connectionType + "\n\n" +

		"\t\t\t\t\tSeguranca da rede: " + connProfile.wifiSecurityType + "\n\n" +

		"\t\t\t\t\tNome da Rede Wi-fi: " + connProfile.essid + "\n\n" +

		"\t\t\t\t\tModo de aquisição IP: " + connProfile.ipMode + "\n\n" +

		"\t\t\t\t\tSenha da Rede: " + connProfile.wifiPassword + "\n\n" +

		"\t\t\t\t\tRede Oculta: " + checkResponse(connProfile.hidden) + "\n")

}

func (connProfile *ConnectionProfile) setConnectionProfile() *ConnectionProfile {
	var wifiInterface, connectionType, wifiSecurityType, essid, ipMode, wifiPassword, hidden, wifiProfileValidation string

	for {
		fmt.Printf("%s", execute("clear"))

		fmt.Printf("Informe sua interface de rede \n")
		fmt.Println(detectNetwork())
		fmt.Scanf("%s", &wifiInterface)

		fmt.Printf("Informe o tipo de conexão \n")
		fmt.Println("[wireless ethernet]")
		fmt.Scanf("%s", &connectionType)

		fmt.Printf("Informe o tipo de segurança wi-fi \n")
		fmt.Println("[wpa wpa2]")
		fmt.Scanf("%s", &wifiSecurityType)

		fmt.Printf("Informe o modo de aquisição de endereço Ip \n")
		fmt.Println("[dhcp]")
		fmt.Scanf("%s", &ipMode)

		fmt.Printf("Informe a identificação da rede \n")
		fmt.Scanf("%s", &essid)

		fmt.Printf("Informe a senha da rede \n")
		fmt.Scanf("%s", &wifiPassword)

		fmt.Printf("A rede está oculta? \n")
		fmt.Println("[sim não]")
		fmt.Scanf("%s", &hidden)

		*connProfile = ConnectionProfile{wifiInterface, connectionType, wifiSecurityType, essid, ipMode, wifiPassword, checkBool(hidden)}

		connProfile.printConnectionProfile()

		fmt.Println("Os dados estao corretos?")
		fmt.Println("[sim não]")
		fmt.Scanf("%s", &wifiProfileValidation)

		if checkBool(wifiProfileValidation) == true {
			break
		}

	}

	return connProfile
}

type Locale struct {
	language       string
	keyboardLayout string
	timezone       string
}

func (locale *Locale) writeLocale() {

	if locale.language == "PTBR" {
		data := []byte("LANG=pt_BR.UTF-8")
		err := ioutil.WriteFile("/etc/locale.conf", data, 0777)
		check(err)
	}

	if locale.keyboardLayout == "br-abnt2" {

		_ = execute("loadkeys br-abnt2")

		data := []byte("KEYMAP=br.abnt2" + "\n" +
			"FONT=lat1-16.psfu.gz")

		err := ioutil.WriteFile("/etc/vconsole.conf", data, 0777)
		check(err)
	}

	if locale.timezone == "Brasília" {
		_ = execute("timedatectl set-ntp true")
		_ = execute("ln -s -f /usr/share/zoneinfo/Brazil/East /etc/localtime")
		_ = execute("hwclock --systohc --utc")

	}
}

func (locale *Locale) setLocale() *Locale {
	var language, keyboardLayout, timezone, localeValidation string

	for {

		fmt.Println("Qual língua você deseja que o sistema possua?")
		fmt.Println("[PTBR]")
		fmt.Scanf("%s", &language)

		fmt.Println("Existe um padrão de teclado definido para sua linguagem. Você desa usá-la?")
		fmt.Println("[br-abnt2]")
		fmt.Scanf("%s", &keyboardLayout)

		fmt.Println("Fuso horário")
		fmt.Println("[Brasília]")
		fmt.Scanf("%s", &timezone)

		*locale = Locale{language, keyboardLayout, timezone}

		locale.printLocale()

		fmt.Println("Os dados estão corretos?")
		fmt.Println("[sim não]")
		fmt.Scanf("%s", &localeValidation)

		if localeValidation == "sim" {
			break
		}

	}

	return locale
}

func (locale *Locale) printLocale() {

	fmt.Println("Os dados selecionados foram:" +
		"Língua : " + locale.language + "\n" +
		"Padrão de teclado : " + locale.keyboardLayout + "\n" +
		"Fuso horário : " + locale.timezone)

}

type Partition struct {
	device         string
	filesystem     string
	partitionTable string
}

func (partition *Partition) setPartition() *Partition {
	var device, filesystem, partitionTable, partitionValidation, overwriteAccept string

	for {
		fmt.Println("Informe em qual dispositivo você deseja criar o particionamento")
		fmt.Println(detectDevice())
		fmt.Scanf("%s", &device)

		fmt.Println("Informe o sistema de arquivos desejado")
		fmt.Println("[f2fs]")
		fmt.Scanf("%s", &filesystem)

		fmt.Println("Informe a tabela de partição")
		fmt.Println("[msdos gpt]")
		fmt.Scanf("%s", &partitionTable)

		*partition = Partition{device, filesystem, partitionTable}

		partition.printPartition()

		fmt.Println("Os dados estão corretos?")
		fmt.Println("[sim não]")
		fmt.Scanf("%s", &partitionValidation)

		if partitionValidation == "sim" {
			break
		}

	}

	fmt.Println("Todos os dados serão apagados da unidade selecionada. Deseja continuar?")
	fmt.Println("[sim não]")
	fmt.Scanf("%s", &overwriteAccept)

	if overwriteAccept == "não" {
		fmt.Println("Você cancelou a instalação")
		os.Exit(0)
	}
	return partition

}

func (partition *Partition) printPartition() {

	fmt.Println("Partição selecionada: " + partition.device + "\n" +
		"Sistema de arquivos:: " + partition.filesystem + "\n" +
		"Tabela de partição : " + partition.partitionTable)

}

func setUefi() bool {
	var efiSupport string
	fmt.Println("Seu computador tem suporte a EFI?")
	fmt.Println("[sim não]")
	fmt.Scanf("%s", &efiSupport)

	return checkBool(efiSupport)

}

func (partition *Partition) writePartitionTable(uefiEnabled bool) {

	if uefiEnabled == true {
		fmt.Println(partition)
		_ = execute("parted -s -a optimal /dev/" + partition.device + " mklabel gpt mkpart ESP 1 500 mkpart primary 500 100% set 1 boot on")
		_ = execute("mkfs.fat -F32 /dev/" + partition.device + "1")
		_ = execute("mkfs.f2fs /dev/" + partition.device + "2")

	} else {

		_ = execute("parted -s -a optimal /dev/" + partition.device + " mkpart primary 0% 100%")
		_ = execute("mkfs.f2fs /dev/" + partition.device + "1")

	}
}

func getUuidPartition(partition string) string {

	cmdOut := execute("blkid -s UUID " + partition)
	r, _ := regexp.Compile(`"([^"]*)"`)
	regex := r.FindStringSubmatch(string(cmdOut))
	return regex[1]

}

func writeBootConfiguration(uuid string) {

	data := []byte("title" + "\t" + "Arch Linux" +
		"linux" + "\t" + "/vmlinuz-linux" +
		"initrd" + "\t" + "/initramfs-linux.img" +
		"options" + "\t" + "root=/dev/disk/by-uuid/" + uuid)
	err := ioutil.WriteFile("/mnt/boot/loader/entries/arch.conf", data, 0777)
	check(err)

}
func writeFstab(fstab []byte) {

	err := ioutil.WriteFile("/mnt/etc/fstab", fstab, 0777)
	check(err)

}

func setHostname() {

	var hostname string

	fmt.Println("Informe o nome desejado para o seu computador")
	fmt.Scanf("%s", hostname)

	data := []byte(hostname)

	err := ioutil.WriteFile("/mnt/etc/hostname", data, 0777)

	check(err)
}

func setPassword(user string) {
	fmt.Println("Digite uma senha para o usuário " + user)
	_ = execute("passwd " + user)

}

func addUser(username string, password string) {
	

	fmt.Println("Digite o nome do usuário")
	fmt.Scanf("%s", username)

	_ = execute("useradd -m -s /bin/bash -G wheel,users,audio,video,input " + username + " -p " + password)

}

func copyBaseConfig() {

	_ = execute("cp /etc/vconsole.conf /mnt/etc/vconsole.conf")
	_ = execute("cp /etc/locale.conf /mnt/etc/locale.conf")
	//_ = execute("cp /etc/netctl/* /mnt/etc/netctl")
//	connProfile.writeWifiConfigToFile("/mnt/etc/netctl")
	_ = execute("ln -s -f /mnt/usr/share/zoneinfo/Brazil/East/ /mnt/etc/localtime")

}

//func start(args ...string) (p *os.Process, err error) {
//	if args[0], err = exec.LookPath(args[0]); err == nil {
//		var procAttr os.ProcAttr
//		procAttr.Files = []*os.File{os.Stdin,
//			os.Stdout, os.Stderr}
//		p, err := os.StartProcess(args[0], args, &procAttr)
//		if err == nil {
//			return p, nil
//		}
//	}
//	return nil, err

//        if proc, err := start("bash", "-c", "mkdir hello; mkdir badabes"); err == nil {
//        fmt.Println("Feliz Natal, ho ho ho")
//               proc.Wait()
//}

func main() {
	var connProfile ConnectionProfile
	var locale Locale
	var partition Partition
	var uefi bool
	var uuid string

	verbose = false

	//teclado

	locale = *locale.setLocale()

	locale.writeLocale()

	//conexão de rede
	connProfile = *connProfile.setConnectionProfile()

	connProfile.writeWifiConfigToFile("/etc/netctl")

	//_ = execute("netctl start " + connProfile.essid) //substitui o wifi-menu

	//particionamento

	uefi = setUefi()

	partition = *partition.setPartition()

	partition.writePartitionTable(uefi)

	_ = execute("mount /dev/" + partition.device + "2 /mnt")
	_ = execute("mkdir -p /mnt/boot")
	_ = execute("mount /dev/" + partition.device + "1 /mnt/boot")

	//	rankMirrors()
	fmt.Println("Instalando base Arch")


	_ = execute("pacstrap /mnt base base-devel")

	
	fmt.Println("Efetuando configura��es adicionais")

	out := execute("genfstab -U /mnt") // > /mnt/etc/fstab")

	writeFstab(out)

	copyBaseConfig()

	
	_=executeInArchChroot("mkinitcpio -p linux")

	_=executeInArchChroot("pacman -S f2fs-tools ntfs-3g dosfstools --noconfirm")

	_=executeInArchChroot("pacman -S intel-ucode --noconfirm")
	
	_=executeInArchChroot("bootctl install")

	fmt.Println(	
	
	uuid = getUuidPartition("/dev/" + partition.device + "2")
	
	writeBootConfiguration(uuid)

	setHostname()

	_ = executeInArchChroot("pacman -S iw wpa_supplicant dialog --noconfirm")

	//		setPassword("root")

	//		addUser()

	//		setPassword(user)

	//	_ = execute ("umount -R /mnt")

	//Pós instalação

	//Drivers intel

	_ = execute("pacman -S 	xf86-video-intel mesa mesa-libgl libva-intel-driver libva")

	_ = execute ("pacman -S pantheon --noconfirm")
		
	fmt.Println ("Instala��o finalizada - Divirta-se :)")
	
}
