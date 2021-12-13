<template>
  <select
    :value="propData"
    v-bind="$attrs"
    class="appearance-none block w-full border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
    @input="updateValue"
  >
    <option value="PRIVATE">🔐 非公開</option>
    <option value="FRIENDS">👯 友人のみ</option>
    <option value="PUBLIC">🔓 公開</option>
  </select>
</template>

<script>
export default {
  props: {
    propData: [String],
  },
  methods: {
    updateValue(event) {
      this.$emit('input', event.target.value)
    },
  },
}
</script>
