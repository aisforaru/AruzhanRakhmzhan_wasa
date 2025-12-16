<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<script>
export default {}
</script>

<template>

	<header class="navbar sticky-top flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="">WASAText</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/home" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								My Conversations
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/me" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
								My Profile
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/search" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
								Search People
							</RouterLink>
						</li>
					</ul>

					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Groups</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink :to="'/groups'" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
								My Groups
							</RouterLink>
						</li>
					</ul>
				</div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style>
body {
  background-color: #1c1c1e !important;
  color: #f2b7c7ff;
  font-family: 'Inter', sans-serif;
}

#sidebarMenu {
  background-color: #2a2a2d !important;
  height: 100vh;
  margin-top: 60px; /* отступ вниз под navbar */
  border-right: 1px solid #3a3a3c;
}

#sidebarMenu .nav-link {
  color: #ffdde6 !important;
  font-size: 15px;
}

#sidebarMenu .nav-link:hover {
  background-color: #3a3a3c !important;
  border-radius: 6px;
}

.sidebar-heading {
  color: #ffdde6 !important;
  font-weight: 600;
  letter-spacing: 0.5px;
}

.login-container {
  background-color: #2f2f33 !important;
  border: 1px solid #3a3a3c !important;
  color: #fff !important;
  border-radius: 12px;
  padding: 32px;
}
.login-container input.form-control {
  background-color: #f6c5f9ff !important;     /* slightly lighter than card (#2f2f33) */
  border: 1px solid #6a6a6d !important;      /* softer border */
  color: #f2f2f7 !important;
}

.login-container input.form-control::placeholder {
  color: #c6c6c8 !important;                 /* soft grey */
}

.login-title {
  color: #fbbfcfff !important;  
  text-align: center;
  font-weight: 700;
  text-shadow:
    0px 0px 2px #6b4f4f,
    0px 0px 4px #6b4f4f;
}

main {
  display: flex;
  justify-content: center;
  align-items: center;
  height: calc(100vh - 60px); /* minus navbar height */
}

.navbar {
  background-color: #2a2a2d !important;
  padding-top: 18px !important;
  padding-bottom: 18px !important;
  border-bottom: 1px solid #3a3a3c;
}

.navbar-brand {
  background-color: transparent !important;
  color: #ffdde6 !important;             /* pastel pink */
  font-weight: 700 !important;
  font-size: 28px !important;            /* bigger title */

  text-shadow:
    0 0 2px #6b4f4f,
    0 0 4px #6b4f4f;                /* brown glow */
}
.navbar .col-md-3,
.navbar .col-lg-2 {
  background-color: transparent !important;
}


.conversation-block,
.list-group-item {
  background-color: #2f2f30ff !important;   /* тёмно-зелёный */
  border: 1px solid #ebb8cdff !important;   /* розовая рамка */
  color: #f2b7c7ff;
}

/* hover */
.conversation-block:hover,
.list-group-item:hover {
  background-color: #3a3a3c !important;
}

/* если внутри есть текст */
.conversation-block p,
.conversation-block span {
  color: #f2b7c7ff;
}

main {
  background-color: transparent !important;
  color: #f2f2f7;
}


.card, .modal-content {
  background-color: #3a3a3c !important;
  border: 1px solid #4a4a4d !important;
  color: #f2f2f7;
  border-radius: 12px;
}


input.form-control {
  background-color: #3a3a3c !important;
  border: 1px solid #4a4a4d !important;
  color: #fff !important;
}

input.form-control::placeholder {
  color: #b3b3b7 !important;
}

.btn-outline-secondary {
  color: #f2b7c7ff !important;
  border-color: #ff779b !important;
  background-color: transparent !important;
}

.btn-outline-secondary:hover {
  background-color: #ff779b !important;
  color: #1c1c1e !important;
}

/* New group */
.btn-outline-primary {
  color: #3a3a3c !important;
  border-color: #ff779b !important;
  background-color: #ff779b !important;
}

.btn-outline-primary:hover {
  background-color: #ff779b !important;
  color: #1c1c1e !important;
}

</style>