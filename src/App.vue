<template>
    <div>
        <button class="button mb-4 is-succes">Carpeta +</button>

        <div class="columns is-multiline">
            <div v-for="image in images" :key="image" class="column is-one-quarter">
                <div class="card">
                    <div  class="card-image selection">
                        <img :src="image" alt="">
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
const buscar_imagenes = require('./utils/lector.js').lector

const fs = window.require('fs');
const configFile = '/home/elias/Documentos/electron-app/configuracion/configuracion.json';
let carpetas = JSON.parse(fs.readFileSync(configFile));

export default {
    name: 'app',
    data () {
        return {
            images: []
        }    
    },
    mounted() {
        this.images = buscar_imagenes(carpetas); // cuando se monta, cargar las imagenes
        console.log(this.images);
    },
    // TODO seguir agregando los metodos para agregar carpetas o para refrescar la pagina
}
</script>
