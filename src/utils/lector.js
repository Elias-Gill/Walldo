
// funcion que busca imagenes dentro de las carpetas configuradas
function lector(){
    // TODO agregar la funcion que lee el archivo de configuracion y despues itera las carpetas
    //se guarda en un array las imagenes encontradas en la carpeta
    let imagenes = importAll(require.context('/home/elias/Imagenes/fondos', true, /\.(jpg$|png$)/));
    return imagenes;
} 

function importAll(r) {
    //retorna todas las imagenes que se encontraron en esa carpeta
    let images = [];
    r.keys().forEach(key => {
        images.push({ pathLong: r(key), pathShort: key })
    });
    return images;
}

exports.lector = lector;
