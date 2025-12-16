<script>
export default {
  data() {
    return {
      username: "",
      errormsg: null,
      loading: false,
      groups: []
    };
  },
  methods: {
    async loadGroups() {
      this.errormsg = null;
      this.loading = true;
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          this.$router.push({ path: "/" });
          return;
        }
        const response = await this.$axios.get("/groups", {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.groups = response.data || [];
      } catch (error) {
        console.error("Error loading groups:", error);
        this.errormsg = "Failed to load groups. Please try again.";
      } finally {
        this.loading = false;
      }
    },
    viewGroup(groupId, groupName) {
        localStorage.setItem("groupName", groupName);
        this.$router.push({
            path: `/groups/${groupId}`
        });
    },
    refresh() {
        this.loadGroups();
    },
    logOut() {
        this.$router.push({ path: "/" });
    },
    newGroup() {
        this.$router.push({ path: "/new-group" });
    }
  },
  mounted() {
    this.username = localStorage.getItem("name") || "Guest";
    this.loadGroups();
  }
};
</script>

<template>
  <div>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="h2">{{ username }}, here are your groups</h1>
      <div class="btn-toolbar mb-2 mb-md-0">
        <div class="btn-group me-2">
          <button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">Refresh</button>
          <button type="button" class="btn btn-sm btn-outline-secondary" @click="logOut">Log Out</button>
        </div>
        <div class="btn-group me-2">
          <button type="button" class="btn btn-sm btn-outline-primary" @click="newGroup">New group</button>
        </div>
      </div>
    </div>
    <ErrorMsg v-if="errormsg" :msg="errormsg" />
    <div>
      <p v-if="loading">Loading...</p>
      <div v-else-if="groups.length === 0">
        <p>No groups found.</p>
      </div>
      <div v-else class="conversations-container">
        <div
          v-for="group in groups"
          :key="group.id"
          class="conversation-block"
          @click="viewGroup(group.id, group.name)"
        >
          <div class="conversation-photo">
            <img
              v-if="group.conversationPhoto.String"
              :src="'data:image/png;base64,' + group.conversationPhoto.String"
              alt="Group Photo"
              class="profile-picture"
            />
          </div>
          <div class="conversation-details">
            <h4>{{ group.name }}</h4>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

h1.h2 {
  color: #f2b7c7ff !important;   /* розовый */
  text-shadow: 0 0 2px #6b4f4f;
}

.d-flex {
  display: flex;
}

.justify-content-between {
  justify-content: space-between;
}

.flex-wrap {
  flex-wrap: wrap;
}

.align-items-center {
  align-items: center;
}

.pt-3 {
  padding-top: 1rem;
}

.pb-2 {
  padding-bottom: 0.5rem;
}

.mb-3 {
  margin-bottom: 1rem;
}

.border-bottom {
  border-bottom: 1px solid #ff779b;
}

.btn-toolbar {
  display: flex;
  flex-wrap: wrap;
}

.btn-group {
  position: relative;
  display: inline-flex;
  vertical-align: middle;
}

.me-2 {
  margin-right: 0.5rem;
}

.mb-2 {
  margin-bottom: 0.5rem;
}

.mb-md-0 {
  margin-bottom: 0;
}


@media (max-width: 600px) {
  .conversation-block {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>