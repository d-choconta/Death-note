# Death Note App ☠📓

Una aplicación Full Stack inspirada en Death Note, desarrollada para Ryuk. Esta herramienta permite registrar personas, causas y detalles de muerte, con condiciones similares a las reglas del cuaderno original. Los datos se almacenan de forma persistente en una base de datos, y el backend está containerizado con Docker. Además, la aplicación incluye pruebas unitarias para asegurar la calidad de su funcionamiento.

---

## 🧠 Requisitos del Proyecto

- Registrar el *nombre completo* (nombres y apellidos) de una persona.
- Registrar la *causa de muerte* hasta 40 segundos después del nombre.
  - Si no se especifica, la muerte será por *ataque al corazón*.
- Si se especifica la causa, se pueden registrar *detalles* en los siguientes 6 minutos y 40 segundos.
  - En este caso, la persona morirá *40 segundos después* de los detalles.
- Requiere *foto del rostro* de la persona. Si no hay foto, no muere.
- Visualización de personas fallecidas en la interfaz.
- Interfaz responsive con un diseño inmersivo y agradable.
- Backend en Go, frontend en React, API REST.
- Pruebas unitarias incluidas.

---

## 🛠 Tecnologías Utilizadas

### Frontend:
- React + TypeScript
- CSS puro
- Vite o Create React App

---

## ✅ Requisitos previos

Antes de empezar, asegúrate de tener instalado:

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Visual Studio Code](https://code.visualstudio.com/)
- [ npm](https://nodejs.org/) (solo si quieres correr el frontend sin Docker)

---
## 🧾 Instrucciones paso a paso

### 1. Clonar el repositorio

1. Abre *Visual Studio Code*.
2. Presiona Ctrl + Ñ para abrir la terminal integrada.
3. En la terminal, elige una carpeta donde quieras guardar el proyecto.

### 2. Abrir Docker
Abre Docker Desktop desde el menú de inicio.

Espera a que el icono de Docker en la bandeja se vuelva verde o diga "Docker is running".Con la terminal de VS Code ya abierta en la carpeta del proyecto, ejecuta: docker-compose up --build

### 3. Ejecutar Programa
Una vez que Docker haya terminado de levantar los servicios; Abre tu navegador y ve al link que aparece en la terminal de VS code: http://localhost:5173/

Ahí verás la interfaz de usuario de la Death Note.

### 4. Detener Programa
Cuando quieras detener la ejecución`Ctrl + C` y docker-compose down esto apagará y eliminará los contenedores.

## 5.Ejecución de Pruebas Unitarias
Una vez dentro de la carpeta donde esta el proyecto, ejecuta el comando eln la terminal de VS code:go test
Esto buscará todos los archivos *_test.go en el directorio actual y ejecutará las funciones de prueba automáticamente.
