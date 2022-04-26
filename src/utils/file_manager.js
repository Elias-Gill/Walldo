const fs = window.require('fs');

// buscar las imagenes por cada carpeta definida en la config
function findAllImages(){ 
    let images = []
    const configFile = readConfigFile();
    if (configFile == -1) { // si el archivo es recien creado, entonces no hace nada
        return [];
    } 
    for (let folder in configFile.folders) {
        console.log(configFile.folders[folder]);
        let aux = imagesOnFolder(configFile.folders[folder]);
        for (let i in aux) {
            images.push(aux[i]);
        }
    }
    return images;
}

// busca imagenes dentro de las carpetas configuradas
function imagesOnFolder(folder){ // TODO  poner try catchs para avisar de errores
    //se guarda en un array las imagenes encontradas en la carpeta
    let imagenes = [];
    fs.readdir(folder, function(err, res){ // res contiene el NOMBRE de la imagen
        for (let i in res) {
            let image = "file://"+folder+res[i] // guardar direccion entera
            imagenes.push(image);
        }
    })
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
        configFile = "/home/elias/Documentos/electron-app/configuracion/configuracion.json"
        configPath = "/home/elias/Documentos/electron-app/configuracion/"
    }
    // si el archivo no existe, crearlo
    if(!fs.existsSync(configFile)){
        fs.mkdir(configPath,{ recursive: true }, (err) => {
            if (err) {
                console.log("No se pudo crear carpeta de config");
            }
        });

        fs.writeFile(configFile, '{}', function(err){
            console.log("NO se pudo crear config: ", err);
        });
        return -1;
    }
    return JSON.parse(fs.readFileSync(configFile)); // retorna la data del archivo de configuracion
}

exports.findAllImages = findAllImages;
exports.readConfigFile = readConfigFile;
