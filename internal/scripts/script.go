package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// TemplateOptions contiene toda la información para crear y configurar la plantilla.
type Script struct {
	Name string
	Url  string
}

// Run_CreateTemplate ejecuta todo el proceso.
// Esta es la función principal que llamarías desde tu TUI.
func (j Script) Execute() error {

	// --- 1. Clonar el repositorio ---
	fmt.Printf("Clonando %s en %s...\n", j.Url, j.Name)
	// Creamos el comando "git clone <url> <carpeta>"
	cmd := exec.Command("git", "clone", j.Url, j.Name)

	time.Sleep(3 * time.Second)
	// Ejecutamos el comando. 'CombinedOutput' captura lo que se imprime
	// en la terminal (stdout y stderr) por si hay un error.
	if output, err := cmd.CombinedOutput(); err != nil {
		// Si hay un error, lo devolvemos junto con la salida de git
		return fmt.Errorf("error al clonar: %v\nSalida de Git: %s", err, string(output))
	}

	// --- 2. Eliminar la carpeta .git ---
	fmt.Println("Repositorio clonado. Eliminando la carpeta .git...")
	gitDir := filepath.Join(j.Name, ".git")

	// Usamos os.RemoveAll para borrar la carpeta .git y todo su contenido.
	// Esto es el equivalente a 'rmdir /s /q' del .bat, pero multiplataforma.
	if err := os.RemoveAll(gitDir); err != nil {
		return fmt.Errorf("error al eliminar .git: %v", err)
	}
	fmt.Println("Carpeta .git eliminada.")

	// // --- 3. Configurar la plantilla (Buscar y Reemplazar) ---
	// if len(opts.Placeholders) > 0 {
	// 	fmt.Println("Configurando la plantilla (reemplazando texto)...")

	// 	// 'filepath.WalkDir' recorre cada archivo y carpeta en el directorio
	// 	err := filepath.WalkDir(j.Name, func(path string, d fs.DirEntry, err error) error {
	// 		if err != nil {
	// 			return err
	// 		}
	// 		// Ignoramos las carpetas
	// 		if d.IsDir() {
	// 			return nil
	// 		}

	// 		// Leemos el contenido del archivo
	// 		content, err := os.ReadFile(path)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		strContent := string(content)
	// 		modified := false

	// 		// Iteramos sobre todos los placeholders que el usuario quiere reemplazar
	// 		for key, value := range opts.Placeholders {
	// 			if strings.Contains(strContent, key) {
	// 				strContent = strings.ReplaceAll(strContent, key, value)
	// 				modified = true
	// 			}
	// 		}

	// 		// Si modificamos el archivo, lo guardamos
	// 		if modified {
	// 			fmt.Printf("Actualizando %s\n", path)
	// 			// Guardamos el archivo con los mismos permisos que tenía
	// 			if err := os.WriteFile(path, []byte(strContent), d.Type().Perm()); err != nil {
	// 				return err
	// 			}
	// 		}
	// 		return nil
	// 	})

	// 	if err != nil {
	// 		return fmt.Errorf("error al configurar la plantilla: %v", err)
	// 	}
	// 	fmt.Println("Configuración completada.")
	// }

	fmt.Println("\n--- ¡Proceso completado! Tu plantilla está lista. ---")
	return nil
}

func InitialScript(name string, url string) IScript {
	return Script{
		Name: name,
		Url:  url,
	}
}
