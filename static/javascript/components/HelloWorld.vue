<template>
    <div class="container mx-auto">
        <h1 class="text-3xl font-bold underline">{{ this.$store.state.count }}</h1>
        <div class="p-2 text-3xl" @click="increase">click</div>
    </div>
</template>
<script>
import { inject } from 'vue';

export default {
    props: ['dataProperty'],
    data() {
        return {
            message: 'Hello, Vue! JS AJA 12358910',
        }
    },
    created() {
        this.$store = inject('store');
    },
    methods: {
        increase: function (event) {
            this.$store.commit('increment')
            // this.$store.commit("addMessage", this.$store.state.count)
            this.$store.commit('setSnackbarMessage', "test ting" );
            this.$store.commit('showMessage');
        }
    }
}
</script>