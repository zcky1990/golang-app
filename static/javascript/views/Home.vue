<template>
    <div class="container mx-auto">
        <snackBar 
            :show="snackBar.show" 
            :message="snackBar.snackbarMessage" 
            position="top" 
            color="green" 
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
                show: true
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
            console.log("emitted Show")
            let snackBar = this.snackBar;
            snackBar.snackbarMessage = message;
            snackBar.show = true;
            setTimeout(function () {
                snackBar.show = false
            }, 1000);
        },
        closeSnackbar() {
            console.log("emitted Hide")
            this.snackBar.show = false;
        }
    }
}
</script>