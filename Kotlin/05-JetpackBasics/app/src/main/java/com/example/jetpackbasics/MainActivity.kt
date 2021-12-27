package com.example.jetpackbasics

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.animation.animateContentSize
import androidx.compose.animation.core.Spring
import androidx.compose.animation.core.animateDpAsState
import androidx.compose.animation.core.spring
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material.*
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.ExpandLess
import androidx.compose.material.icons.filled.ExpandMore
import androidx.compose.runtime.*
import androidx.compose.runtime.saveable.rememberSaveable
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
    val tempString = ("Composem ipsum color sit lazy, " +
            "padding theme elit, sed do bouncy. ").repeat(4)

    Surface(
        color = MaterialTheme.colors.primary,
        modifier = Modifier
            .padding(vertical = 5.dp, horizontal = 8.dp)
            .fillMaxWidth()
    ) {
        Row(
            modifier = Modifier.padding(24.dp)
                .animateContentSize(
                    spring(
                        dampingRatio = Spring.DampingRatioMediumBouncy,
                        stiffness = Spring.StiffnessLow
                    )
                )
        ) {
            Column(modifier = Modifier.weight(1f)) {
                Text(text = "Hello,")
                Text(text = name, style=MaterialTheme.typography.h4)
                if (isExpanded) Text(text = tempString)
            }

            IconButton(onClick = { isExpanded = !isExpanded }) {
                Icon(
                    imageVector = if (isExpanded) Icons.Filled.ExpandLess else Icons.Filled.ExpandMore,
                    "아이콘"
                )
            }
        }
    }
}


@Composable
fun MyApp() {
    // it은 인덱스
    val names = List(1000) { "$it" }

    LazyColumn (modifier = Modifier.padding(vertical=5.dp)) {
        items(items=names) { name ->
            Greeting(name)
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
    var shouldShowOnboarding by rememberSaveable {mutableStateOf(true)}
    JetpackBasicsTheme {
        if (shouldShowOnboarding) {
            OnboardingScreen({shouldShowOnboarding = false})
        } else {
            MyApp()
        }
    }
}