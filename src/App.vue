<template>
    <div>
        <button class="button mb-4 is-succes">Carpeta +</button>

        <div class="columns is-multiline">
            <div v-for="image in images" :key="image" class="column is-one-quarter">
                <div class="card">
                    <div  class="card-image selection">
                        <img :src="image.pathLong" alt="">
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
const buscar_imagenes = require('./utils/lector.js').lector
export default {
    name: 'app',
    data () {
        return {
            contador: 1,
            images: []
        }    
    },
    mounted() {
        this.images = buscar_imagenes(); // cuando se monta, cargar las imagenes
    },
    // TODO seguir agregando los metodos para agregar carpetas o para refrescar la pagina
}
</script>
