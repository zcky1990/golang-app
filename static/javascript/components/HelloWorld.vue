<template>
    <div class="container mx-auto">
        <h1 class="text-3xl font-bold underline">{{ this.$store.state.count }}</h1>
        <div class="p-2 text-3xl" @click="increase">click</div>
    </div>
</template>
<script>
import { inject } from 'vue';

export default {
    created() {
        this.$store = inject('store');
    },
    methods: {
        increase: function (event) {
            this.$store.commit('increment')
            this.$emit("showSnakeBarCallback","Error", "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam ea quo unde vel adipisci blanditiis voluptates eum. Nam, cum minima?", "error");
        }
    }
}
</script>