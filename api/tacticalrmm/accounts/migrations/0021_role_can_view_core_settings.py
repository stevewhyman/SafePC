# Generated by Django 3.2.4 on 2021-06-17 04:29

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('accounts', '0020_role_can_manage_roles'),
    ]

    operations = [
        migrations.AddField(
            model_name='role',
            name='can_view_core_settings',
            field=models.BooleanField(default=False),
        ),
    ]
