<template>
    <div class="container mx-auto">
        <snackBar 
            :show="snackBar.show" 
            :message="snackBar.snackbarMessage" 
            position="top" 
            color="green" 
            timeout="10000"
            @showSnakeBar="showSnackbar"
            @closeSnakeBar="closeSnackbar" 
        />
        <contents 
            :dataProperty="data" 
            @closeSnakeBar="closeSnackbar" 
            @showSnakeBar="showSnackbar"
        />
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
    components: {
        contents,
        showCount,
        snackBar
    },
    data(){
        return {
            snackBar:{
                snackbarMessage: '',
                show: false
            }
        }
    },
    created() {
        this.$store = inject('store');
    },
    computed: {

    },
    methods: {
        showSnackbar(message) {
            let snackBar = this.snackBar;
            snackBar.snackbarMessage = message;
            snackBar.show = true;
        },
        closeSnackbar() {
            this.snackBar.show = false;
        }
    }
}
</script>