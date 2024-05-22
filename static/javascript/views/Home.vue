<template>
    <div class="container mx-auto">
        <contents :dataProperty="data"/>
        <router-view />
    </div>
</template>
<script>
import contents from "./../components/HelloWorld.vue"

export default {
    name: 'Home',
    props: ['data'],
    components: {
        contents,
    }
}
</script>