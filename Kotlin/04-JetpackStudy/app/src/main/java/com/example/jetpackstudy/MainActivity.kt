package com.example.jetpackstudy

import android.content.res.Configuration
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.Image
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Surface
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.jetpackstudy.ui.theme.JetpackStudyTheme

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            JetpackStudyTheme {
                DefaultPreview()
            }
        }
    }
}

data class Message(
    val author: String,
    val msg: String
)

@Composable
fun MessageCard(msg: Message) {
    Row (modifier = Modifier.padding(all = 10.dp)){
        Image(
            painter = painterResource(R.drawable.profile_picture),
            contentDescription = "기본 사진",
            modifier = Modifier
                .size(40.dp)
                .clip(CircleShape)
                .border(width = 1.dp, color = MaterialTheme.colors.primary, CircleShape)
        )

        Spacer(modifier = Modifier.width(8.dp))

        Column {
            Row {
                Text(
                    text = "Author: ",
                    color = MaterialTheme.colors.secondary,
                    style = MaterialTheme.typography.subtitle1
                )
                Text(
                    text = msg.author,
                    style = MaterialTheme.typography.subtitle1
                )
            }

            Spacer(modifier = Modifier.width(8.dp))

            Surface(shape = MaterialTheme.shapes.medium, elevation = 2.dp) {
                Text(
                    text = "Message: ${msg.msg}",
                    style = MaterialTheme.typography.body1
                )
            }
        }
    }
}

@Preview(showBackground = true)
@Preview(
    uiMode = Configuration.UI_MODE_NIGHT_YES,
    showBackground = true,
    name = "Dark Mode"
)
@Composable
fun DefaultPreview() {
    val message = Message("원석", "메시지")

    JetpackStudyTheme {
        MessageCard(msg = message)
    }
}