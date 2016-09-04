package models

type Partition struct {
	device         string
	filesystem     string
	partitionTable string
}

//func (partition *Partition) setPartition() *Partition {
//	var device, filesystem, partitionTable, partitionValidation, overwriteAccept string

//	for {
//		fmt.Println("Informe em qual dispositivo você deseja criar o particionamento")
//		fmt.Println(detectDevice())
//		fmt.Scanf("%s", &device)

//		fmt.Println("Informe o sistema de arquivos desejado")
//		fmt.Println("[f2fs]")
//		fmt.Scanf("%s", &filesystem)

//		fmt.Println("Informe a tabela de partição")
//		fmt.Println("[gpt]")
//		fmt.Scanf("%s", &partitionTable)

//		*partition = Partition{device, filesystem, partitionTable}

//		partition.printPartition()

//		fmt.Println("Os dados estão corretos?")
//		fmt.Println("[sim não]")
//		fmt.Scanf("%s", &partitionValidation)

//		if partitionValidation == "sim" {
//			break
//		}

//	}

//	fmt.Println("Todos os dados serão apagados da unidade selecionada. Deseja continuar?")
//	fmt.Println("[sim não]")
//	fmt.Scanf("%s", &overwriteAccept)

//	if overwriteAccept == "não" {
//		fmt.Println("Você cancelou a instalação")
//		os.Exit(0)
//	}
//	return partition

//}
