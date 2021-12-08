<template>
  <div>
    <Nav />
    <header class="bg-white shadow">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold text-gray-900">設定</h1>
      </div>
    </header>
    <main>
      <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div class="px-4 py-6 sm:px-0">
          <div class="border-2 border-gray-200 rounded h-96">
            <form class="px-8 pt-6 pb-8 mb-4">
              <div
                v-for="(error, index) in errors"
                :key="index"
                class="mb-4 text-red-600"
              >
                {{ error }}
              </div>

              <div class="mb-4">
                <label
                  class="block text-gray-700 text-sm font-bold mb-2"
                  for="name"
                >
                  名前（表示用）
                </label>
                <input
                  id="name"
                  v-model="user.name"
                  class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                  type="text"
                  placeholder="名無しのクライマー"
                  required
                />
              </div>

              <div class="flex items-center justify-between">
                <button
                  class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                  type="button"
                  @click="onSubmit"
                >
                  更新
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import myself from '~/apollo/myself'

export default {
  apollo: {},
  data() {
    return {
      submitting: false,
      errors: [],
      user: {},
    }
  },
  head() {
    return {
      title: '設定',
    }
  },
  computed: {},
  watch: {},
  beforeMount() {},
  mounted() {},
  async created() {
    try {
      await this.$apollo
        .query({
          query: myself,
          variables() {
            return {}
          },
        })
        .then(({ data }) => {
          const user = data.myself
          this.user = user
        })
    } catch (e) {
      console.log(e)
      // this.$router.push('/logout')
    }
  },
  methods: {
    onSubmit(e) {
      e.preventDefault()
      this.errors = []
      this.submitting = true

      if (this.name !== null) {
        this.name = this.name.trim()
      }

      if (!this.name) {
        this.errors.push('名前を入力してください。')
      }

      if (this.errors.length > 0) {
        return false
      }

      try {
        this.submitting = false
      } catch (e) {
        console.log(e)
        this.submitting = false
        this.errors.push(e)

        window.scrollTo({
          top: 0,
          behavior: 'smooth',
        })
      }
    },
  },
}
</script>

<style></style>
