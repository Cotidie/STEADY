package com.example.jetpackstudy

import android.content.res.Configuration
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.animation.animateColorAsState
import androidx.compose.animation.animateContentSize
import androidx.compose.foundation.Image
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Surface
import androidx.compose.material.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.jetpackstudy.ui.theme.JetpackStudyTheme

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            JetpackStudyTheme {
                Conversation(SampleData.conversationSample)
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

        // mutableState(false)를 기억하는 변수
        // 단순 boolean으로 하지 않는 이유는?
        var isExpanded by remember { mutableStateOf(false) }
        val surfaceColor: Color = if (isExpanded) MaterialTheme.colors.primary else MaterialTheme.colors.surface

        Column (modifier = Modifier.clickable { isExpanded = !isExpanded }){
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

            Surface(
                color = surfaceColor,
                shape = MaterialTheme.shapes.medium,
                elevation = 2.dp,
                modifier = Modifier.animateContentSize()
            ) {
                Text(
                    text = "Message: ${msg.msg}",
                    style = MaterialTheme.typography.body1,
                    maxLines = if (isExpanded) Int.MAX_VALUE else 1
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

@Composable
fun Conversation(msgs: List<Message>) {
    LazyColumn {
        items(msgs) { msg ->
            MessageCard(msg)
        }
    }
}

@Preview(
    showBackground = true,
    name = "Conversation"
)
@Composable
fun PreviewConversation() {
    JetpackStudyTheme {
        Conversation(SampleData.conversationSample)
    }
}