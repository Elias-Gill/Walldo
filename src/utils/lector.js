//const fs = require('fs');
//console.log(fs);

//const direccion = '../configuracion/configuracion.json'  // cambiar si quieres mover la config

// retorna las imagenes encontradas dentro de las carpetas en el 
// archivo de configuracion
/*function buscar_imagenes(){
    const configuracion = fs.readFileSync(direccion); // archivo de configuracion
    const data = JSON.parse(configuracion);

    // cargar las carpetas
    let direcciones = [];
    for (const x in data) {
        direcciones.push(data[x]);
    }

    //cargar las imagenes que estan en las carpetas
    let imagenes = [];
    for (let dir of direcciones) {
        let files = fs.readdirSync(dir) // leer las direcciones de la carpeta
        for (let x of files) {
            if (x.includes('jpg') || x.includes('png')){
                imagenes.push(x);
            }
        }
    }

    console.log(imagenes);
    return imagenes;
}*/

//exports.buscar_imagenes = buscar_imagenes;
