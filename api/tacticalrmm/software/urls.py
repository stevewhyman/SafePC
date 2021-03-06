from django.urls import path

from . import views

urlpatterns = [
    path("chocos/", views.chocos),
    path("install/", views.install),
    path("installed/<int:pk>/", views.get_installed),
    path("refresh/<int:pk>/", views.refresh_installed),
]
