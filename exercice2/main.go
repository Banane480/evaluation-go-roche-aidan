package main

import "fmt"

const ( // Codes couleur ANSI pour le terminal
	Reset = "\033[0m"
	Gras  = "\033[1m"
	Cyan  = "\033[36m"
	Rouge = "\033[31m"
	Vert  = "\033[32m"
)

func main() {
	fmt.Println(Cyan + Gras + "=== GESTIONNAIRE DE NOTES ===" + Reset)
	fmt.Println()

	var nb int
	fmt.Print("Combien de notes voulez-vous saisir ? ")
	fmt.Scan(&nb)

	if nb <= 0 { // Arrêt du programme en cas de saisie invalide
		fmt.Println(Rouge + Gras + "Nombre de notes invalide." + Reset)
		return
	}

	var notes []int // Liste (slice) d'entiers pour stocker les notes
	somme := 0

	for i := 1; i <= nb; i++ { // Boucle pour saisir chaque note, de 1 jusqu'à nb
		var note int
		for { // Boucle de vérification qui s'assure que la note est entre 0 et 20
			fmt.Printf("Note %d : ", i)
			fmt.Scan(&note)

			if note >= 0 && note <= 20 {
				break // Note valide : on sort de la boucle de vérification
			}
			fmt.Println(Rouge + "Note invalide ! Une note doit être comprise entre 0 et 20." + Reset)

		}
		notes = append(notes, note) // On ajoute la note à la liste
		somme += note               // On ajoute la note à la somme
	}

	moyenne := calculerMoyenne(somme, nb)
	max := trouverMaximum(notes)
	min := trouverMinimum(notes)

	//Résulats
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
	return float64(somme) / float64(nombre) // Division en float64 pour obtenir un résultat décimal précis
}

func trouverMaximum(notes []int) int {
	max := notes[0] // On prend la première valeur comme maximum initial
	for _, note := range notes {
		if note > max { // Si la note est plus grande, on met à jour le maximum
			max = note
		}
	}
	return max
}

func trouverMinimum(notes []int) int {
	min := notes[0] // On prend la première valeur comme minimum initial
	for _, note := range notes {
		if note < min { // Si la note est plus petite, on met à jour le minimum
			min = note
		}
	}
	return min
}

func afficherResultat(moyenne float64) { // Affichage du résultat selon la moyenne
	if moyenne >= 10 {
		fmt.Println(Vert + Gras + "Étudiant admis !" + Reset)
	} else {
		fmt.Println(Rouge + Gras + "Non admis" + Reset)
	}
}

// BONUS
func compterNotesAuDessusDeLaMoyenne(notes []int, moyenne float64) int {
	compteur := 0
	for _, note := range notes {
		if float64(note) >= moyenne { // On compare chaque note à la moyenne
			compteur++
		}
	}
	return compteur
}

func afficherAppreciation(moyenne float64) { // Affichage de l'appréciation selon la moyenne
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
}
