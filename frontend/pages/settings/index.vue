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
          <div class="border-2 border-gray-200 rounded">
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

              <div class="flex flex-wrap">
                <label
                  class="block w-full text-gray-700 text-sm font-bold mb-2"
                  for="height"
                >
                  🚻 性別
                </label>
                <div class="w-16 mb-6">
                  <select
                    id="gender"
                    v-model="user.gender"
                    class="appearance-none block w-full border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                  >
                    <option value="MALE">🚹 男</option>
                    <option value="FEMALE">🚺 女</option>
                  </select>
                </div>
                <div class="w-36 px-3 mb-6">
                  <div class="relative">
                    <DisplayFlag :prop-data="user.genderDisplay" />
                  </div>
                </div>
              </div>

              <div class="flex flex-wrap">
                <label
                  class="block w-full text-gray-700 text-sm font-bold mb-2"
                  for="height"
                >
                  ↕️ 身長
                </label>
                <div class="w-16 mb-6">
                  <input
                    id="height"
                    v-model="user.height"
                    class="appearance-none block w-full border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                    type="text"
                    placeholder="160"
                    required
                  />
                </div>
                <div class="w-36 px-3 mb-6">
                  <div class="relative">
                    <DisplayFlag :prop-data="user.heightDisplay" />
                  </div>
                </div>
              </div>

              <div class="flex flex-wrap">
                <label
                  class="block w-full text-gray-700 text-sm font-bold mb-2"
                  for="wingspan"
                >
                  ↔️ リーチ
                </label>
                <div class="w-16 mb-6">
                  <input
                    id="wingspan"
                    v-model="user.wingspan"
                    class="appearance-none block w-full border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                    type="text"
                    placeholder="160"
                    required
                  />
                </div>
                <div class="w-36 px-3 mb-6">
                  <div class="relative">
                    <DisplayFlag :prop-data="user.wingspanDisplay" />
                  </div>
                </div>
              </div>

              <div class="flex flex-wrap">
                <label
                  class="block w-full text-gray-700 text-sm font-bold mb-2"
                  for="weight"
                >
                  ⚖️ 体重
                </label>
                <div class="w-16 mb-6">
                  <input
                    id="weight"
                    v-model="user.weight"
                    class="appearance-none block w-full border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                    type="text"
                    placeholder="160"
                    required
                  />
                </div>
                <div class="w-36 px-3 mb-6">
                  <div class="relative">
                    <DisplayFlag :prop-data="user.weightDisplay" />
                  </div>
                </div>
              </div>

              <div class="flex flex-wrap">
                <label
                  class="block w-full text-gray-700 text-sm font-bold mb-2"
                  for="birthday"
                >
                  🎂 生年月日
                </label>
                <div class="w-42 mb-6">
                  <select>
                    <option value="1970">1970</option>
                  </select>
                  <select>
                    <option value="01">01</option>
                  </select>
                  <select>
                    <option value="01">01</option>
                  </select>
                </div>
                <div class="w-36 md:px-3 mb-6">
                  <div class="relative">
                    <DisplayFlag :prop-data="user.birthdayDisplay" />
                  </div>
                </div>
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

      if (this.user.name !== null) {
        // this.user.name = this.user.name.trim()
      }

      if (!this.user.name) {
        this.errors.push('名前を入力してください。')
      }

      if (this.user.gender !== null) {
        // this.user.gender = this.user.gender.trim()
      }

      if (!this.user.gender) {
        this.errors.push('性別を入力してください。')
      }

      if (this.user.height !== null) {
        // this.user.height = this.user.height.trim()
      }

      if (!this.user.height) {
        this.errors.push('身長を入力してください。')
      }

      if (this.user.wingspan !== null) {
        // this.user.wingspan = this.user.wingspan.trim()
      }

      if (!this.user.wingspan) {
        this.errors.push('リーチを入力してください。')
      }

      if (this.user.birthday !== null) {
        // this.user.birthday = this.user.birthday.trim()
      }

      if (!this.user.birthday) {
        // this.errors.push('生年月日を入力してください。')
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
