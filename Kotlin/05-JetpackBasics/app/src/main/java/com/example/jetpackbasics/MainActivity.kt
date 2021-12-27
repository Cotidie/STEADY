package com.example.jetpackbasics

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.*
import androidx.compose.material.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.jetpackbasics.ui.theme.JetpackBasicsTheme

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            JetpackBasicsTheme {
                OnboardingPreview()
            }
        }
    }
}

@Composable
fun Greeting(name: String) {
    var isExpanded by remember {mutableStateOf(false)}
    val extraPadding = if (isExpanded) 48.dp else 0.dp

    Surface(
        color = MaterialTheme.colors.primary,
        modifier = Modifier
            .padding(vertical = 5.dp, horizontal = 8.dp)
            .fillMaxWidth()
    ) {
        Row(modifier = Modifier.padding(24.dp)) {
            Column(modifier = Modifier
                .weight(1f)
                .padding(bottom = extraPadding)) {
                Text(text = "Hello,")
                Text(text = name)
            }

            OutlinedButton(onClick = {
                isExpanded = !isExpanded
            }) {
                Text(
                    text = if (isExpanded) "Show Less" else "Show More"
                )
            }
        }
    }
}


@Composable
fun MyApp() {
    val names = listOf("Andrew", "Finn")

    Column (modifier = Modifier.padding(vertical=5.dp)){
        for (name in names) {
            Greeting(name = name)
        }
    }
}

@Composable
fun OnboardingScreen(onContinueClicked: () -> Unit) {
    Surface {
        Column(
            modifier = Modifier.fillMaxSize(),
            verticalArrangement = Arrangement.Center,
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Text("Welcome to the Basics Codelab!")
            Button(
                modifier = Modifier.padding(vertical = 24.dp),
                onClick = onContinueClicked
            ) {
                Text("Continue")
            }
        }
    }
}

@Preview(showBackground = true, widthDp = 320, heightDp = 320)
@Composable
fun OnboardingPreview() {
    var shouldShowOnboarding by remember {mutableStateOf(true)}
    JetpackBasicsTheme {
        if (shouldShowOnboarding) {
            OnboardingScreen({shouldShowOnboarding = false})
        } else {
            MyApp()
        }
    }
}