<script>
export default {
  data() {
    localStorage.clear();
    return {
      errormsg: null,
      name: "", 
      profile: {
        id: "",
        name: "",
      },
    };
  },
  methods: {
    async loadDefaultPhoto() {
      try {
          const response = await fetch('/nopfp.jpg');
          const blob = await response.blob();
          const reader = new FileReader();
          return new Promise((resolve, reject) => {
              reader.onload = () => {
                  const base64String = reader.result.toString().split(',')[1];
                  resolve(base64String);
              };
              reader.onerror = reject;
              reader.readAsDataURL(blob);
          });
      } catch (error) {
          console.error('Error loading default photo:', error);
          return null;
      }
  },
    blobToBase64(blob) {
      return new Promise((resolve) => {
        const reader = new FileReader();
        reader.onloadend = () => resolve(reader.result);
        reader.readAsDataURL(blob);
      });
    },
    async doLogin() {
      if (this.name.trim() === "") {
        this.errormsg = "Name cannot be empty.";
        return;
      }
      try {
        const photoData = await this.loadDefaultPhoto();
        const response = await this.$axios.post("/session", {
          name: this.name,
          photo: photoData,
        }, {
          headers: {
            'Content-Type': 'application/json'
          }
        });
        if (response.data.identifier) {
          this.profile.id = response.data.identifier;
          this.profile.name = this.name; 
        } else {
          throw new Error("Unexpected server response. Missing 'identifier'.");
        }
        localStorage.setItem("token", this.profile.id);
        localStorage.setItem("name", this.profile.name);
        this.$router.push({ path: "/home" });
      } catch (e) {
        if (e.response && e.response.status === 400) {
          this.errormsg =
            "Form error, please check all fields and try again.";
        } else if (e.response && e.response.status === 500) {
          this.errormsg =
            "An internal error occurred. Please try again later.";
        } else {
          this.errormsg = e.toString();
        }
      }
    },
  },
};
</script>

<template>
  <div class="login-container shadow-sm">
    <h2 class="login-title">Welcome to WASAText</h2>
    <div class="input-group">
      <input
        type="text"
        id="name"
        v-model="name"
        class="login-input"
        placeholder="Insert your name to log in WASAText."
      />
      <button class="login-button" type="button" @click="doLogin">Login</button>
    </div>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </div>
</template>

<style scoped>
.login-container {
  max-width: 420px;
  margin: 140px auto;
  text-align: center;
  padding: 28px;
  border-radius: 14px;

  /* dark grey card matching your theme */
  background-color: #2f2f33;
  border: 1px solid #3a3a3c;

  box-shadow: 0 6px 14px rgba(0, 0, 0, 0.35);
}

.login-title {
  font-size: 26px;
  font-weight: 700;
  margin-bottom: 22px;
  color: #fcdfe7ff;

  text-shadow:
    0 0 2px #6b4f4f,
    0 0 3px #6b4f4f;
}

.input-group {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.login-input {
  padding: 12px;
  width: 100%;
  margin-bottom: 14px;
  background-color: #e9c1ccff;
  border: 1px solid #6a6a6d;
  border-radius: 6px;
  color: #504447ff;
  font-size: 15px;
}

.login-button {
  padding: 10px 15px;
  background-color: #ff8eacff;
  color: white;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  font-size: 16px;
}

.login-input::placeholder {
  color: #5f555fff;
}

.login-button:hover {
  background-color: #ff779bff;
}
</style>
