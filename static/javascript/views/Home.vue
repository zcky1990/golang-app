<template>
    <div>
        <snackBar :show="snackBar.show" :message="snackBar.snackbarMessage" position="top" color="green" timeout=10000
            :type="snackBar.type" :title="snackBar.title" @showSnakeBar="showSnackbar" @closeSnakeBar="closeSnackbar" />
        <contents :dataProperty="data" @closeSnakeBar="closeSnackbar" @showSnakeBar="showSnackbar" />
        <showCount />
        <login />
        <signUp />
    </div>
</template>
<script>
import { inject } from 'vue';

import snackBar from "./../components/shared/Snackbar.vue";
import contents from "./../components/HelloWorld.vue";
import showCount from "../components/ShowCount.vue";
import login from "../components/Login.vue"
import signUp from "../components/SignUp.vue"
import { info } from 'autoprefixer';

export default {
    name: 'Home',
    components: {
        contents,
        showCount,
        snackBar,
        login,
        signUp
    },
    data() {
        return {
            snackBar: {
                snackbarMessage: 'test',
                show: true,
                type: 'info',
                title: ''
            }
        }
    },
    created() {
        this.$store = inject('store');
    },
    computed: {

    },
    methods: {
        showSnackbar(title, message, type) {
            let snackBar = this.snackBar;
            snackBar.title = title;
            if (type === 'error') {
                snackBar.type = type;
            } else {
                snackBar.type = "info";
            }
            snackBar.snackbarMessage = message;
            snackBar.show = true;
        },
        closeSnackbar() {
            this.snackBar.show = false;
        }
    }
}
</script>