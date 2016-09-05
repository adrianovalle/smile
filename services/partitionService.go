



//func getUuidPartition(partition string) string {

//	cmdOut := execute("blkid -s UUID " + partition)
//	r, _ := regexp.Compile(`"([^"]*)"`)
//	regex := r.FindStringSubmatch(string(cmdOut))
//	return regex[1]

//}

//func writeBootConfiguration(uuid string) {

//	//	_ = execute("mkdir /mnt/boot/loader")
//	//	_ = execute("mkdir /mnt/boot/loader/entries")

//	data := []byte("title" + "\t" + "Arch Linux" + "\n" +
//		"linux" + "\t" + "/vmlinuz-linux" + "\n" +
//		"initrd" + "\t" + "/initramfs-linux.img" + "\n" +
//		"options" + "\t" + "root=/dev/disk/by-uuid/" + uuid)
//	err := ioutil.WriteFile("/mnt/boot/loader/entries/arch.conf", data, 0666)
//	check(err)

//}
//func writeFstab(fstab []byte) {

//	err := ioutil.WriteFile("/mnt/etc/fstab", fstab, 0777)
//	check(err)

//}

//func (partition *Partition) writePartitionTable(uefiEnabled bool) {

//	if uefiEnabled == true {
//		fmt.Println(partition)
//		_ = execute("parted -s -a optimal /dev/" + partition.device + " mklabel gpt mkpart ESP 1 500 mkpart primary 500 100% set 1 boot on")
//		_ = execute("mkfs.fat -F32 /dev/" + partition.device + "1")
//		_ = execute("mkfs.f2fs /dev/" + partition.device + "2")

//	} else {

//		_ = execute("parted -s -a optimal /dev/" + partition.device + " mkpart primary 0% 100%")
//		_ = execute("mkfs.f2fs /dev/" + partition.device + "1")

//	}
//}

//func setUefi() bool {
//	var efiSupport string
//	fmt.Println("Seu computador tem suporte a EFI?")
//	fmt.Println("[sim não]")
//	fmt.Scanf("%s", &efiSupport)

//	return checkBool(efiSupport)

//}

//func (partition *Partition) printPartition() {

//	fmt.Println("Partição selecionada: " + partition.device + "\n" +
//		"Sistema de arquivos:: " + partition.filesystem + "\n" +
//		"Tabela de partição : " + partition.partitionTable)

//}