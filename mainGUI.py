import os
import tkinter as tk
from tkinter import scrolledtext
import subprocess

# Establecer el directorio de trabajo al mismo que el script
ruta_script = os.path.dirname(os.path.abspath(__file__))
os.chdir(ruta_script)


def guardar_texto():
    codigo = text_area.get("1.0", tk.END)
    with open("entrada.txt", "w") as archivo:
        archivo.write(codigo)

def ejecutar_programa_go():
    subprocess.run(["go", "run", "main.go"])

def mostrar_resultados():
    with open("salida.txt", "r") as archivo:
        resultado = archivo.read()
    resultado_text.config(state=tk.NORMAL)
    resultado_text.delete("1.0", tk.END)
    resultado_text.insert(tk.END, resultado)
    resultado_text.config(state=tk.DISABLED)

# Configuración de la ventana principal
ventana = tk.Tk()
ventana.title("Analizador Léxico")

# Etiqueta de instrucciones
instrucciones_label = tk.Label(ventana, text="Ingrese las líneas de código:")
instrucciones_label.pack()

# Área de texto para ingresar código
text_area = scrolledtext.ScrolledText(ventana, wrap=tk.WORD, width=40, height=10)
text_area.pack(padx=10, pady=10)  # Agrega márgenes a la izquierda y derecha


# ... (configuración de widgets como en el ejemplo anterior)

# Botón para guardar el texto en un archivo
guardar_button = tk.Button(ventana, text="Guardar Texto", command=guardar_texto)
guardar_button.pack(pady=5)  # Añade espacio vertical

# Botón para ejecutar el programa en Go
ejecutar_button = tk.Button(ventana, text="Ejecutar en Go", command=ejecutar_programa_go)
ejecutar_button.pack(pady=5)  # Añade espacio vertical

# Botón para mostrar los resultados
mostrar_button = tk.Button(ventana, text="Mostrar Resultados", command=mostrar_resultados)
mostrar_button.pack(pady=5)  # Añade espacio vertical

# Recuadro de resultados
resultado_text = scrolledtext.ScrolledText(ventana, wrap=tk.WORD, width=40, height=10, state=tk.DISABLED)
resultado_text.pack(pady=10)  # Agrega margen hacia abajo

# Inicia el bucle de la aplicación
ventana.mainloop()