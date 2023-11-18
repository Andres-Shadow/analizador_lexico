import os
import tkinter as tk
from tkinter import scrolledtext
from tkinter import ttk
from tkinter import messagebox
import subprocess

# Establecer el directorio de trabajo al mismo que el script
ruta_script = os.path.dirname(os.path.abspath(__file__))
os.chdir(ruta_script)

def guardar_texto():
    codigo = text_area.get("1.0", tk.END)
    with open("entrada.txt", "w") as archivo:
        archivo.write(codigo)
    messagebox.showinfo("Guardado", "Texto guardado exitosamente.")

def ejecutar_programa_go():
    try:
        subprocess.run(["go", "run", "main.go"])
        mostrar_resultados()
    except FileNotFoundError:
        messagebox.showerror("Error", "El ejecutable de Go no se encuentra. Asegúrate de tener Go instalado.")

def mostrar_resultados():
    try:
        with open("salida.txt", "r") as archivo:
            resultado = archivo.read()
        resultado_text.config(state=tk.NORMAL)
        resultado_text.delete("1.0", tk.END)
        resultado_text.insert(tk.END, resultado)
        resultado_text.config(state=tk.DISABLED)
    except FileNotFoundError:
        messagebox.showwarning("Advertencia", "El archivo de resultados no se encuentra. Ejecuta el programa en Go primero.")

# Configuración de la ventana principal
ventana = tk.Tk()
ventana.title("Analizador Léxico")
ventana.geometry("600x500")  # Tamaño predeterminado
ventana.resizable(False, False)  # Ventana no redimensionable

# Estilo de los botones
style = ttk.Style()
style.configure("TButton", padding=10, relief="flat", background="#4CAF50", foreground="black", font=("Arial", 12))

# Etiqueta de instrucciones
instrucciones_label = tk.Label(ventana, text="Ingrese las líneas de código:", font=("Arial", 14))
instrucciones_label.pack(pady=10)

# Área de texto para ingresar código
text_area = scrolledtext.ScrolledText(ventana, wrap=tk.WORD, width=50, height=10, font=("Courier New", 12))
text_area.pack(padx=20, pady=10)  # Agrega márgenes a la izquierda y derecha

# Frame para contener los botones
botones_frame = ttk.Frame(ventana)
botones_frame.pack()

# Botón para guardar el texto en un archivo
guardar_button = ttk.Button(botones_frame, text="Guardar Texto", command=guardar_texto)
guardar_button.pack(side=tk.LEFT, padx=10)  # Alinea a la izquierda con un espacio

# Botón para ejecutar el programa en Go
ejecutar_button = ttk.Button(botones_frame, text="Ejecutar en Go", command=ejecutar_programa_go)
ejecutar_button.pack(side=tk.LEFT, padx=10)  # Alinea a la izquierda con un espacio

# Botón para mostrar los resultados
mostrar_button = ttk.Button(botones_frame, text="Mostrar Resultados", command=mostrar_resultados)
mostrar_button.pack(side=tk.LEFT, padx=10)  # Alinea a la izquierda con un espacio

# Recuadro de resultados
resultado_text = scrolledtext.ScrolledText(ventana, wrap=tk.WORD, width=50, height=10, state=tk.DISABLED, font=("Courier New", 12))
resultado_text.pack(pady=20)  # Agrega margen hacia abajo

# Inicia el bucle de la aplicación
ventana.mainloop()
