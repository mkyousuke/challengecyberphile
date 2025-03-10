package main
 
import (
    "fmt"
    "github.com/phpdave11/gofpdf"
    "io"
    "log"
    "net/http"
    "os"
)
 
func generatePDF() {
    // Créer un nouveau PDF
    pdf := gofpdf.New("P", "mm", "A4", "")
 
    // Ajouter une page
    pdf.AddPage()
 
    // Ajouter la police DejaVuSans
    pdf.AddUTF8Font("DejaVu", "", "fonts/DejaVuSans.ttf")
 
    // Définir la police
    pdf.SetFont("DejaVu", "", 14)
 
    // Ajouter un titre
    pdf.Cell(200, 10, "Rapport d'Audit de Securite")
    pdf.Ln(12) // Nouvelle ligne
 
    // Ajouter le contenu du rapport
    pdf.MultiCell(0, 10, "Voici un exemple de rapport généré automatiquement.\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam. Sed nisi.", "", "L", false)
 
    // Sauvegarder le fichier PDF
    err := pdf.OutputFileAndClose("rapport_audit_securite.pdf")
    if err != nil {
        log.Fatalf("Erreur lors de la sauvegarde du PDF: %s\n", err)
    }
 
    fmt.Println("PDF généré avec succès!")
}
 
func handleDownload(w http.ResponseWriter, r *http.Request) {
    // Vérifier si la méthode est GET
    if r.Method == http.MethodGet {
        // Créer le PDF
        generatePDF()
 
        // Ouvrir le fichier PDF
        file, err := os.Open("rapport_audit_securite.pdf")
        if err != nil {
            http.Error(w, "Erreur d'ouverture du fichier PDF", http.StatusInternalServerError)
            return
        }
        defer file.Close()
 
        // Spécifier le type de contenu
        w.Header().Set("Content-Type", "application/pdf")
        w.Header().Set("Content-Disposition", "attachment; filename=rapport_audit_securite.pdf")
 
        // Copier le contenu du fichier dans la réponse
        _, err = io.Copy(w, file)
        if err != nil {
            http.Error(w, "Erreur lors de l'envoi du fichier PDF", http.StatusInternalServerError)
        }
    } else {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
    }
}
 
func handleHome(w http.ResponseWriter, r *http.Request) {
    // Page HTML simple avec lien pour télécharger le PDF
    htmlContent := `
        <!DOCTYPE html>
        <html lang="fr">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Rapport d'Audit de Sécurité</title>
        </head>
        <body>
            <h1>Bienvenue sur le rapport d'audit de sécurité</h1>
            <p><a href="/download">Cliquez ici pour télécharger le rapport d'audit en PDF</a></p>
        </body>
        </html>
    `
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(htmlContent))
}
 
func main() {
    // Définir les routes
    http.HandleFunc("/", handleHome)         // Page d'accueil
    http.HandleFunc("/download", handleDownload) // Téléchargement du PDF
 
    // Lancer le serveur HTTP
    fmt.Println("Serveur lancé sur : http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
