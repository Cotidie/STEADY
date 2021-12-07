package com.example.quizapp

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import com.example.quizapp.databinding.ActivityQuizQuestionsBinding

class QuizQuestionsActivity : AppCompatActivity() {
    lateinit var binding: ActivityQuizQuestionsBinding
    private val questionList = Constants.getQeustions(5)

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityQuizQuestionsBinding.inflate(layoutInflater)

        setQuestion(0)
        setContentView(binding.root)
    }

    // idx번 퀴즈를 설정한다,
    private fun setQuestion(idx: Int) {
        val question = questionList.get(idx)
        binding.ivFlag.setImageResource(question.image)
        binding.progressBar.progress = idx
        binding.tvProgress.setText("${idx+1}/${binding.progressBar.max}")
        binding.tvQuestion.setText(question.question)
        binding.tvOptionOne.setText(question.optionOne)
        binding.tvOptionTwo.setText(question.optionTwo)
        binding.tvOptionThree.setText(question.optionThree)
        binding.tvOptionFour.setText(question.optionFour)
    }
}