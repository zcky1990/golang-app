<template>
    <div>
        <snackBar :show="snackBar.show" :message="snackBar.snackbarMessage" position="top" color="green" :timeout=1000
            :type="snackBar.type" :title="snackBar.title" @showSnakeBarCallback="showSnackbar"
            @closeSnackebarCallback="closeSnackbar" />
        <countDown position="center" targetDate="2025-03-25" />
        <dropdown v-model:value="dropdown.value" :dropdown-label="dropdown.dropdownLabel" :dropdown-items="dropdown.dropdownItem" v-model:show-error="dropdown.error" />
        <selectComp v-model:value="dropdown.value" :select-label="dropdown.dropdownLabel" :select-items="dropdown.dropdownItem" v-model:show-error="dropdown.error"/>
        <dividerComp :message="divider.message" :message-position="divider.messagePosition" :color="divider.color"
            :message-transform="divider.messageTransform" />
        <prokes />
        <notificationWithAction :show-dialog="dialogPopUp.showNotification" :title-dialog="dialogPopUp.title"
            :message-dialog="dialogPopUp.message" :type-dialog="dialogPopUp.type" :button-dialog="dialogPopUp.button"
            @submitCallback="closeDialog" @cancelCallback="closeDialog" />
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
import dividerComp from "./../components/Divider.vue";
import countDown from "./../components/CountDown.vue"
import dropdown from "./../components/Dropdown.vue"
import selectComp from "./../components/Select.vue"
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
        dropdown,
        selectComp,
        countDown,
        dividerComp,
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
                show: false,
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
            },
            divider: {
                message: "test",
                messagePosition: "center",
                messageTransform: "uppercase",
                color: "red",
            },
            dropdown: {
                dropdownLabel: "Dropdown Label Example",
                value: '',
                dropdownItem: [{
                    key: "Dropdown 1",
                    value: "1"
                },
                {
                    key: "Dropdown 2",
                    value: "2"
                }],
                error: true
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
        },

    }
}
</script>