<template>
    <div>
        <snackBar :show="snackBar.show" :message="snackBar.snackbarMessage" position="top" color="green" :timeout=1000
            :type="snackBar.type" :title="snackBar.title" @showSnakeBarCallback="showSnackbar" @closeSnackebarCallback="closeSnackbar" />
        <countDown position="center" targetDate="2025-03-25"/>
        <arrumCard />
        <prokes />
        <notificationWithAction 
            :showDialog="dialogPopUp.showNotification" 
            :titleDialog="dialogPopUp.title"
            :messageDialog="dialogPopUp.message" 
            :typeDialog="dialogPopUp.type"
            :buttonDialog="dialogPopUp.button"
            @submitCallback="closeDialog"
            @cancelCallback="closeDialog"
        />
        <forbidden />
        <contents :dataProperty="data" @closeSnackebarCallback="closeSnackbar" @showSnakeBarCallback="showSnackbar" />
        <showCount />
        <login />
        <signUp />
    </div>
</template>
<script>
import { inject } from 'vue';

import snackBar from "./../components/shared/Snackbar.vue";
import envelope from "./../components/Envelope.vue";
import countDown from "./../components/CountDown.vue"
import arrumCard from "./../components/ArrumCard.vue"
import forbidden from "./../components/Forbidden.vue";
import contents from "./../components/HelloWorld.vue";
import showCount from "../components/ShowCount.vue";
import notificationWithAction from "../components/NotificationWithAction.vue"
import prokes from "../components/Prokes.vue";
import login from "../components/Login.vue"
import signUp from "../components/SignUp.vue"
import { info } from 'autoprefixer';

export default {
    name: 'Home',
    components: {
        contents,
        showCount,
        countDown,
        arrumCard,
        snackBar,
        notificationWithAction,
        prokes,
        login,
        signUp,
        forbidden,
        envelope
    },
    data() {
        return {
            snackBar: {
                snackbarMessage: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam ea quo unde vel adipisci blanditiis voluptates eum. Nam, cum minima?',
                show: true,
                type: 'info',
                title: ''
            },
            dialogPopUp: {
                showNotification: false,
                title: 'Notification',
                message: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam ea quo unde vel adipisci blanditiis voluptates eum. Nam, cum minima?',
                type: 'notification',
                button: {
                    submit: 'Submit',
                    cancel: 'Cancel'
                }
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
            this.dialogPopUp.showNotification = true
        },
        closeSnackbar() {
            this.snackBar.show = false;
        },
        showDialog() {
            this.dialogPopUp.showNotification = true
        },
        closeDialog() {
            this.dialogPopUp.showNotification = false
        }

    }
}
</script>