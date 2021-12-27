package com.example.jetpackstudy

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Surface
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.tooling.preview.Preview
import com.example.jetpackstudy.ui.theme.JetpackStudyTheme

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            DefaultPreview()
        }
    }
}

data class Message(
    val author: String,
    val msg: String
)

@Composable
fun MessageCard(msg: Message) {
    Row {
        Image(
            painter = painterResource(R.drawable.profile_picture),
            contentDescription = "기본 사진"
        )

        Column {
            Text("Author: ${msg.author}")
            Text("Message: ${msg.msg}")
        }
    }
}

@Preview(showBackground = true)
@Composable
fun DefaultPreview() {
    val message = Message("원석", "메시지")
    MessageCard(msg = message)
}