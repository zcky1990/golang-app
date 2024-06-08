<template>
    <div class="container mx-auto">
        <contents :dataProperty="data"/>
        <showCount/>
    </div>
</template>
<script>
import contents from "./../components/HelloWorld.vue"
import showCount from "../components/ShowCount.vue";

export default {
    name: 'Home',
    props: ['data'],
    components: {
        contents,
        showCount
    }
}
</script>