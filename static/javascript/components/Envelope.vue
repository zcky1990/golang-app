<template>
    <div v-if="show" class="grid content-center h-screen gap-2 px-4 bg-white">
        <div class="flex flex-col justify-center gap-4 mb-10 text-center">
            <p class="text-3xl tracking-widest text-gray-500 uppercase">The wedding of</p>
            <p class="text-2xl tracking-widest text-gray-500 uppercase">Bride</p>
            <p class="text-xl tracking-widest text-gray-500 uppercase">&</p>
            <p class="text-2xl tracking-widest text-gray-500 uppercase">Groom</p>
            <p class="tracking-widest text-gray-500 uppercase text-md">15.01.2024</p>

            <div class="flex flex-col gap-4 mt-6">
                <p class="tracking-widest text-gray-500 text-md">Kepada yth:</p>
                <p class="tracking-widest text-gray-500 text-md">Bpk/ibu/saudara/i</p>
                <p class="text-xl tracking-widest text-gray-500">SMTI-05</p>
            </div>

            <div class="flex flex-row justify-center">
                <a @click="closeEnvelope"
                    class="inline-block w-2/4 px-8 py-3 text-sm font-medium text-indigo-600 transition border border-current rounded-full hover:scale-110 hover:shadow-xl focus:outline-none focus:ring active:text-indigo-500"
                    href="#">
                    Open Invitation
                </a>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            show: true,
        }
    },
    methods:{
        closeEnvelope(){
            this.show = false;
        }
    }
}
</script>