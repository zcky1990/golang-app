<template>
    <div class="container mx-auto">
        <snackBar :show="showSnackbar" :message="snackbarMessage" position="top" color="green" @close="closeSnackbar" />
        <contents :dataProperty="data"/>
        <showCount/>
    </div>
</template>
<script>
import { inject } from 'vue';

import snackBar from "./../components/shared/Snackbar.vue";
import contents from "./../components/HelloWorld.vue";
import showCount from "../components/ShowCount.vue";

export default {
    name: 'Home',
    props: ['data'],
    components: {
        contents,
        showCount,
        snackBar
    },
    created() {
        this.$store = inject('store');
    },
    computed: {
        showSnackbar() {
            return this.$store.state.snackbar.show;
        },
        snackbarMessage() {
            return this.$store.state.snackbar.message;
        }
    },
    methods: {
        showMessage() {
            this.$store.dispatch('snackbar/showSnackbar', 'Hello, Snackbar!');
        },
        closeSnackbar() {
            this.$store.commit('snackbar/hideMessage');
        }
  }
}
</script>