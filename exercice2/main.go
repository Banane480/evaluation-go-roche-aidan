package main

import "fmt"

const (
	Reset = "\033[0m"
	Gras  = "\033[1m"
	Cyan  = "\033[36m"
	Rouge = "\033[31m"
	Vert  = "\033[32m"
) //codes couleur golang (j'ai rajouté en plus)

func main() {
	fmt.Println(Cyan + Gras + "=== GESTIONNAIRE DE NOTES ===" + Reset)
	fmt.Println()

	var nb int
	fmt.Print("Combien de notes voulez-vous saisir ? ")
	fmt.Scan(&nb)

	if nb <= 0 {
		fmt.Println(Rouge + Gras + "Nombre de notes invalide." + Reset)
		return //arrêt du programme.
	}

	var notes []int // J'ai utilisé une liste d'int, je trouve ca assez pratique + l'instruction append est facile.
	somme := 0      // somme des notes pour la moyenne par la suite.

	for i := 1; i <= nb; i++ { // Cette boucle permet de récupérer le nombre de notes qu'on veut, elle va de 1 jusqu'a nb.
		var note int
		for {
			fmt.Printf("Note %d : ", i) //affiche : note1=... note2=... etc...
			fmt.Scan(&note)

			if note >= 0 && note <= 20 {
				break //relance la boucle avec le même i si la note est pas entre 0 ou 20
			}
			fmt.Println(Rouge + "Note invalide ! Une note doit être comprise entre 0 et 20." + Reset)
		} // Boucle infinie qui vérifie que la note est entre 0 et 20, sinon il la redemande.
		notes = append(notes, note) // On ajoute la note a la liste.
		somme += note               // On ajoute la note a la somme.
	}

	moyenne := calculerMoyenne(somme, nb)
	max := trouverMaximum(notes)
	min := trouverMinimum(notes)

	fmt.Println()
	fmt.Println(Cyan + Gras + "=== RÉSULTATS ===" + Reset)
	fmt.Println()
	fmt.Printf("Moyenne : %.2f\n", moyenne)
	fmt.Printf("Note maximale : %d\n", max)
	fmt.Printf("Note minimale : %d\n", min)
	fmt.Println()

	afficherResultat(moyenne)

	//BONUS
	fmt.Printf("Nombre de notes au-dessus de la moyenne : %d\n", compterNotesAuDessusDeLaMoyenne(notes, moyenne))
	afficherAppreciation(moyenne)
}

func calculerMoyenne(somme int, nombre int) float64 {
	return float64(somme) / float64(nombre) //division en float54 pour éviter les overflow + comaptibilité avec les décimaux
}

func trouverMaximum(notes []int) int {
	max := notes[0] // on prend la première valeur comme maximum par défaut.
	for _, note := range notes {
		if note > max { //si la prochaine est plus grande que la défaut alors on la remplace.
			max = note
		}
	}
	return max
}

func trouverMinimum(notes []int) int {
	min := notes[0] // on prend la première valeur comme minimum par défaut.
	for _, note := range notes {
		if note < min { //si la prochaine est plus petite que la défaut alors on la remplace.
			min = note
		}
	}
	return min
}

func afficherResultat(moyenne float64) {
	if moyenne >= 10 {
		fmt.Println(Vert + Gras + "Étudiant admis !" + Reset)
	} else {
		fmt.Println(Rouge + Gras + "Non admis" + Reset)
	}
} // simple comparaison.

// BONUS
func compterNotesAuDessusDeLaMoyenne(notes []int, moyenne float64) int {
	compteur := 0
	for _, note := range notes { //notes = notre liste de notes
		if float64(note) >= moyenne { // On compare chaque note à la moyenne
			compteur++
		}
	}
	return compteur
}

func afficherAppreciation(moyenne float64) {
	if moyenne < 10 {
		fmt.Println(Rouge + "Appréciation : Insuffisant" + Reset)
	} else if moyenne < 12 {
		fmt.Println("Appréciation : Passable")
	} else if moyenne < 14 {
		fmt.Println("Appréciation : Assez bien")
	} else if moyenne < 16 {
		fmt.Println(Vert + "Appréciation : Bien" + Reset)
	} else {
		fmt.Println(Vert + Gras + "Appréciation : Très bien" + Reset)
	}
} //les commentaire dépendants de la moyenne + codes couleurs
