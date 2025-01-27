document.addEventListener('DOMContentLoaded', () => {
    const loginToggle = document.getElementById('login-toggle');
    const signupToggle = document.getElementById('signup-toggle');
    const loginForm = document.getElementById('login-form');
    const signupForm = document.getElementById('signup-form');

    // Form toggle logic
    loginToggle.addEventListener('click', () => {
        loginToggle.classList.add('active');
        signupToggle.classList.remove('active');
        loginForm.classList.add('active');
        signupForm.classList.remove('active');
    });

    signupToggle.addEventListener('click', () => {
        signupToggle.classList.add('active');
        loginToggle.classList.remove('active');
        signupForm.classList.add('active');
        loginForm.classList.remove('active');
    });

    // Enhanced form validation and submission
   signupForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const formData = new FormData(signupForm);
    
    // Log the form data to console for debugging
    for (let pair of formData.entries()) {
        console.log(pair[0] + ': ' + pair[1]);
    }

    try {
        const response = await fetch('/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: new URLSearchParams(formData).toString()
        });

        // Log the response status
        console.log('Response status:', response.status);

        if (response.redirected) {
            showSuccess('Registration successful! Redirecting to login...');
            setTimeout(() => {
                window.location.href = response.url;
            }, 1500);
            return;
        }

        const text = await response.text();
        console.log('Response text:', text);

        if (!response.ok) {
            throw new Error(text || 'Registration failed');
        }

    } catch (error) {
        console.error('Error:', error);
        showError(error.message || 'An error occurred during registration');
    }
});

    // Error/Success message handling
    function showError(message) {
        const errorDiv = document.createElement('div');
        errorDiv.className = 'error-message';
        errorDiv.style.cssText = `
            color: #e74c3c;
            background: #ffd7d7;
            padding: 10px;
            border-radius: 5px;
            margin-bottom: 15px;
            font-size: 14px;
        `;
        errorDiv.textContent = message;

        // Remove any existing error messages
        const existingError = signupForm.querySelector('.error-message');
        if (existingError) {
            existingError.remove();
        }

        // Insert error message at the top of the form
        signupForm.insertBefore(errorDiv, signupForm.firstChild);

        // Animate the error message
        gsap.from(errorDiv, {
            opacity: 0,
            y: -20,
            duration: 0.3,
            ease: 'power2.out'
        });

        // Remove error message after 5 seconds
        setTimeout(() => {
            gsap.to(errorDiv, {
                opacity: 0,
                y: -20,
                duration: 0.3,
                ease: 'power2.in',
                onComplete: () => errorDiv.remove()
            });
        }, 5000);
    }

    function showSuccess(message) {
        const successDiv = document.createElement('div');
        successDiv.className = 'success-message';
        successDiv.style.cssText = `
            color: #27ae60;
            background: #d4ffd4;
            padding: 10px;
            border-radius: 5px;
            margin-bottom: 15px;
            font-size: 14px;
        `;
        successDiv.textContent = message;

        signupForm.insertBefore(successDiv, signupForm.firstChild);

        gsap.from(successDiv, {
            opacity: 0,
            y: -20,
            duration: 0.3,
            ease: 'power2.out'
        });
    }

    // GSAP Animations
    gsap.from('.form-container', {
        opacity: 0,
        y: 50,
        duration: 0.8,
        ease: 'power3.out'
    });

    gsap.from('form input', {
        opacity: 0,
        x: -50,
        stagger: 0.2,
        duration: 0.6,
        ease: 'power2.out'
    });

    gsap.from('.site-footer', {
        opacity: 0,
        y: 50,
        duration: 0.8,
        delay: 0.6,
        ease: 'power3.out'
    });
});