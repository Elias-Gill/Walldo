const fs = window.require('fs');
// busca imagenes dentro de las carpetas configuradas
function lector(carpetas){
    //se guarda en un array las imagenes encontradas en la carpeta
    let imagenes = [];
    for (let carpeta in carpetas){
        fs.readdir(carpetas[carpeta], function(err, res){
            if (err) {
                console.log(err);
            } else {
                for (let i in res) {
                    let aux = "file://"+carpetas[carpeta]+res[i]
                    imagenes.push(aux);
                }
            }
        });
    }
    console.log(imagenes);
    return imagenes;
} 
exports.lector = lector;
