const fs = window.require('fs');
// abrir el explorador
function openExplorer(current){
    let path = '';
    let command = ''
    if(process.platform == 'win32'){
        command = 'explorer.exe'
    } else {
        command = 'pcmanfm'
    }
    return path // TODO terminar esta funcion para que sirva
}

// buscar las imagenes por cada carpeta definida en la config
function findAllImages(){ 
    let images = []
    const configFile = readConfigFile();
    if (configFile == -1) {
        return [];
    } 
    for (folder in configFile.folders) {
        let aux = imageFinder(configFile.folders[folder])
        for (let image in aux) {
            images.push(image);
        }
    }
    return images;
}

// busca imagenes dentro de las carpetas configuradas
function imageFinder(folder){ // TODO  poner try catchs para avisar de errores
    //se guarda en un array las imagenes encontradas en la carpeta
    let imagenes = [];
    for (let image in folder){
        fs.readdir(folder[image], function(err, res){ // res contiene el NOMBRE de la imagen
            if (err) {
                console.log(err);
            } else {
                for (let i in res) {
                    let aux = "file://"+folder[image]+res[i] // guardar direccion entera
                    imagenes.push(aux);
                }
            }
        });
    }
    return imagenes;
} 

// leer el archivo de configuracion ()
function readConfigFile(){
    let configFile = '';
    let configPath = '';
    if (process.platform == 'win32'){
        configFile = "C:/Program Files/Walldo/config/config.json"
        configPath = "C:/Program Files/Walldo/config/"
    } else {
        configFile = "/home/elias/Documents/electron-app/configuracion/configuracion.json"
        configPath = "/home/elias/Documents/electron-app/configuracion/"
    }
    // si el archivo no existe, crearlo
    if(!fs.existsSync(configFile)){
        fs.mkdir(configPath);
        fs.writeFile(configFile, '', (err)=>{
            new Notification({ title: "Permission denied", body: "Try running the app with root privilegies(run as administrator)" }).show()
        });
        return -1;
    };
    return JSON.parse(fs.readFileSync(configFile));
}

exports.findAllImages = findAllImages;
exports.readConfigFile = readConfigFile;
exports.openExplorer = openExplorer;
