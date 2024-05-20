<template>
    <div class="container">
        <contents />
        <router-view />
    </div>
</template>
<script>
import contents from "./../components/HelloWorld.vue"

export default {
    name: 'App',
    components: {
        contents,
    }
}
</script>