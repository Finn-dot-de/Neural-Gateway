import os
import platform
import subprocess

def print_word_files(folder_path):
    """
    Findet alle Word-Dateien in einem Ordner und druckt sie mit dem Standarddrucker.
    Funktioniert unter Windows mit Microsoft Word und unter macOS mit lpr.
    """
    system_platform = platform.system()
    
    for file_name in os.listdir(folder_path):
        if file_name.endswith(".docx") or file_name.endswith(".doc"):
            file_path = os.path.join(folder_path, file_name)
            print(f"Drucke Datei: ({file_path})")
            
            try:
                if system_platform == "Windows":
                    # Windows: Microsoft Word zum Drucken nutzen
                    subprocess.run(["cmd", "/c", "start", "winword", "/q", "/n", "/t", file_path, "/mFilePrintDefault", "/mDuplex=1", "/mFileExit"], shell=True)
                elif system_platform == "Darwin":
                    # macOS: lpr Befehl mit Duplex-Option verwenden
                    subprocess.run(["lpr", "-o", "sides=two-sided-long-edge", file_path])
                else:
                    print("Drucken wird auf diesem Betriebssystem nicht unterstützt.")
            except Exception as e:
                print(f"Fehler beim Drucken von {file_name}: {e}")

if __name__ == "__main__":
    ordner = "./"
    if os.path.exists(ordner):
        print_word_files(ordner)
    else:
        print("Fehler: Der Ordner existiert nicht!")
