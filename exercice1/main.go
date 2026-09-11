package main

import "fmt"

const ( //codes couleurs golang
	Reset = "\033[0m"
	Gras  = "\033[1m"
	Cyan  = "\033[36m"
	Rouge = "\033[31m"
	Vert  = "\033[32m"
)

func main() {
	for {
		afficherMenu()

		var choix int
		fmt.Println()
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix) //récupère le choix

		if choix == 0 {
			fmt.Println("Aurevoir !")
			break
		} // Option quitter du menu

		if choix < 1 || choix > 4 {
			fmt.Println(Rouge + "Choix invalide" + Reset)
			fmt.Println()
			continue
		} // Vérifie si le choix est valide (compris entre 1 et 4)

		fmt.Println()          // Saut de ligne pour la lisibilité
		afficherBoisson(choix) // Affiche la boisson choisie

		prix := obtenirPrix(choix) // Récupère le prix de la boisson
		fmt.Printf("Prix : %d €", prix)
		fmt.Println()
		fmt.Println(Vert + "-----------------------" + Reset)

		var montant int
		fmt.Println() //saut de ligne après les ---
		fmt.Print("Montant donné :")
		fmt.Scan(&montant)

		if montant < prix {
			manque := prix - montant
			fmt.Printf(Rouge+Gras+"Il manque %d €\n"+Reset, manque)
			fmt.Println()
			continue
		} // Vérifie si le montant est suffisant

		rendu := montant - prix // Calcul du rendu
		fmt.Println()
		fmt.Println(Vert + "Merci pour votre achat !" + Reset)
		fmt.Printf("Votre monnaie : %d €\n", rendu)
	}
}

func afficherMenu() {
	fmt.Println(Cyan + Gras + "=== DISTRIBUTEUR ===" + Reset)
	fmt.Println()
	fmt.Println("1 - Eau       : 1 €")
	fmt.Println("2 - Soda      : 2 €")
	fmt.Println("3 - Café      : 2 €")
	fmt.Println("4 - Chocolat  : 3 €")
	fmt.Println("0 - Quitter")
	fmt.Println(Vert + "-----------------------" + Reset)
} // Affiche le menu

func obtenirPrix(choix int) int {
	switch choix {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 2
	case 4:
		return 3
	default:
		return 0
	}
} // switch, similaire à if... else if...

func afficherBoisson(choix int) {
	switch choix {
	case 1:
		fmt.Println("Vous avez choisi : Eau")
	case 2:
		fmt.Println("Vous avez choisi : Soda")
	case 3:
		fmt.Println("Vous avez choisi : Café")
	case 4:
		fmt.Println("Vous avez choisi : Chocolat")
	default:
		fmt.Println(Rouge + "Choix invalide" + Reset)
	}
} // Affiche le nom de la boisson selon le choix
