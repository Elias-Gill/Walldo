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
const buscarImagenes = require('./utils/file_manager.js').findAllImages
export default {
    name: 'app',
    data () {
        return {
            images: []
        }    
    },
    mounted() {
        this.images = buscarImagenes(); // cuando se monta, cargar las imagenes

        /* console.log(this.images); */
    },
    // TODO seguir agregando los metodos para agregar carpetas o para refrescar la pagina
}
</script>
