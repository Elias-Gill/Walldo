<template>
    <div>
        <button class="button mb-4 is-succes">Carpeta +</button>
        <div class="columns is-multiline">

            <div v-for="image in images[0]" :key="image" class="column is-one-quarter">
                <div class="card">
                    <div  class="card-image selection">
                        <img :src="image.pathLong" alt="">
                    </div>
                </div>
            </div>

            <div v-for="image in images[1]" :key="image" class="column is-one-quarter">
                <div class="card">
                    <div  class="card-image selection">
                        <img :src="image.pathLong" alt="">
                    </div>
                </div>
            </div>

            <div v-for="image in images[2]" :key="image" class="column is-one-quarter">
                <div class="card">
                    <div  class="card-image selection">
                        <img :src="image.pathLong" alt="">
                    </div>
                </div>
            </div>

            <div v-for="image in images[3]" :key="image" class="column is-one-quarter">
                <div class="card">
                    <div  class="card-image selection">
                        <img :src="image.pathLong" alt="">
                    </div>
                </div>
            </div>

        </div>
    </div>
</template>

<style>
.is-focus{
    overflow: none;
    border: none;
}

.selection{
    transition: all 2s;
}
</style>

<script>
export default {
    name: 'app',
    data () {
        return {
            contador: 1,
            images: [[],[],[],[]]
        }    
    },
    mounted() {
        // TODO anadir metodo que retorne el array de paths desde archivo de configuraciones
        this.images = this.importAll(require.context('/home/elias/Imagenes/fondos', true, /\.jpg$/));
    },

    methods: {
        importAll(r) {
            let images = [[],[],[],[]]
            let contador = 0;
            r.keys().forEach(key => {
                if(contador == 4){
                    contador = 0;
                }
                images[contador].push({ pathLong: r(key), pathShort: key })
                contador++;
            });
            console.log(this);
            return images
        },
    },
}
</script>
